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

func verifyAgeCommitment(age int, ageCommitment AgeCommitment) bool {
	ageBigInt := new(big.Int).SetInt64(int64(age))
	return zksigma.Open(PatientLedgerCurve, ageBigInt, &ageCommitment.randomValue, ageCommitment.commitment)
}
