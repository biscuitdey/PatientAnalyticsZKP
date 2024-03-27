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

/*
Exponential map: v -> g^v
Scalar multiplication: v -> v * G
Age Commitment = g^v * h^r = vG + rH
Secret Key (sk) = s
Public Key (pk) = p = h^s = sH
Token Commitment = r * p = r * s * H = h^(r * s)
*/

func verifyAgeCommitment(expectedAge int, computedAge *AgeCommitment, pk zksigma.ECPoint, sk *big.Int) bool {
	expectedAgeBitInt := new(big.Int).SetInt64(int64(expectedAge))

	//vG = totalAge * G                               ==> g^v
	vg := PatientLedgerCurve.Mult(PatientLedgerCurve.G, expectedAgeBitInt)
	//Additive inverse of vG = -vG                    ==> multiplicative inverse of g^v = 1/g^v
	minusgv := PatientLedgerCurve.Neg(vg)
	//Commitment - vG = vG + rH - vG = rH             ==> Commitment/g^v = h^v
	T := PatientLedgerCurve.Add(computedAge.commitment, minusgv)

	//Equivalence proof
	/*
		EXPONENTIAL
			BasePoint1 = h^r
			Result1 = h^(r * s)
			BasePoint2 = h
			Result2 = h^(s)
			Equivalence:  log_h^r (h^rs) = log_h(p)
	*/

	/*
		SCALAR MULTIPLICATION
			BasePoint1 = rH
			Result1 = s * rH
			BasePoint2 = H
			Result2 = s * H
			Equal scalar value = s
	*/
	equivalenceProof, _ := zksigma.NewEquivalenceProof(PatientLedgerCurve, T, computedAge.randomValueToken, PatientLedgerCurve.H, pk, sk)
	isVerified, _ := equivalenceProof.Verify(PatientLedgerCurve, T, computedAge.randomValueToken, PatientLedgerCurve.H, pk)

	return isVerified
}
