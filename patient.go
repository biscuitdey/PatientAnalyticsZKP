package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

type AgeCommitment struct {
	commitment  zksigma.ECPoint
	randomValue big.Int
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

func MakeEncryptedPatient(disease string, age int) *EncryptedPatient {

	diseaseZKP := generateDiseaseProof(disease)

	ageCommitment := generateAgeCommitment(age)

	encryptedPatient := &EncryptedPatient{
		disease: *diseaseZKP,
		age:     *ageCommitment,
	}
	return encryptedPatient
}

func generateDiseaseProof(disease string) *ZeroKnowledgeProof {
	return generateProof(disease)
}

func generateAgeCommitment(age int) *AgeCommitment {
	ageBigInt := new(big.Int).SetInt64(int64(age))
	commitment, randomValue, _ := zksigma.PedCommit(PatientLedgerCurve, ageBigInt)
	ageCommitment := &AgeCommitment{
		commitment:  commitment,
		randomValue: *randomValue,
	}

	return ageCommitment
}

func verifyAgeCommitment(expectedAge int, computedAgeCommitment zksigma.ECPoint) bool {
	expectedAgeBitInt := new(big.Int).SetInt64(int64(expectedAge))
	expectedAgeCommitment, _, _ := zksigma.PedCommit(PatientLedgerCurve, expectedAgeBitInt)

	if expectedAgeCommitment.X.Cmp(computedAgeCommitment.X) == 0 && expectedAgeCommitment.Y.Cmp(computedAgeCommitment.Y) == 0 {
		return true
	}
	return false
}
