package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

type AgeCommitment struct {
	commitment       zksigma.ECPoint
	randomValueToken zksigma.ECPoint
}

type Patient struct {
	disease string
	age     int
}
type EncryptedPatient struct {
	id      int
	disease ZeroKnowledgeProof
	age     AgeCommitment
}

func (ledger *Ledger) MakeEncryptedPatient(disease string, age int) *EncryptedPatient {

	diseaseZKP := generateDiseaseProof(disease)

	ageCommitment := ledger.generateAgeCommitment(age)

	encryptedPatient := &EncryptedPatient{
		disease: *diseaseZKP,
		age:     *ageCommitment,
	}
	return encryptedPatient
}

func generateDiseaseProof(disease string) *ZeroKnowledgeProof {
	return generateProof(disease)
}

func (ledger *Ledger) generateAgeCommitment(age int) *AgeCommitment {
	ageBigInt := new(big.Int).SetInt64(int64(age))
	commitment, randomValue, _ := zksigma.PedCommit(PatientLedgerCurve, ageBigInt)
	randomValueToken := zksigma.CommitR(PatientLedgerCurve, ledger.Key.pk, randomValue)

	ageCommitment := &AgeCommitment{
		commitment:       commitment,
		randomValueToken: randomValueToken,
	}

	return ageCommitment
}

func verifyAgeCommitment(expectedAge int, computedAge *AgeCommitment, pk zksigma.ECPoint, sk *big.Int) bool {
	expectedAgeBitInt := new(big.Int).SetInt64(int64(expectedAge))

	gv := PatientLedgerCurve.Neg(PatientLedgerCurve.Mult(PatientLedgerCurve.G, expectedAgeBitInt)) // 1 / g^\sum{v_i}
	T := PatientLedgerCurve.Add(computedAge.commitment, gv)

	equivalenceProof, _ := zksigma.NewEquivalenceProof(PatientLedgerCurve, T, computedAge.randomValueToken, PatientLedgerCurve.H, pk, sk)
	isVerified, _ := equivalenceProof.Verify(PatientLedgerCurve, T, computedAge.randomValueToken, PatientLedgerCurve.H, pk)

	return isVerified
}
