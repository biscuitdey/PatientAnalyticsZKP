package main

import (
	"crypto/sha256"
	"math/big"

	"github.com/mit-dci/zksigma"
)

func hash(data string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte("string"))
	md := hasher.Sum(nil)

	return md
}

func hashToBigInt(data []byte) *big.Int {
	dataBigInt := new(big.Int).SetBytes(data)
	return dataBigInt
}

func calculateAverageAge(patients []Patient) int {
	if len(patients) == 0 {
		return 0
	}

	var totalAge int
	for _, patient := range patients {
		totalAge += patient.age
	}

	return totalAge / len(patients)
}

func calculateEncryptedAverageAge(encryptedPatients []EncryptedPatient) zksigma.ECPoint {
	return computeAverage(encryptedPatients)
}

func filterPatients(disease string, patients []Patient) (filteredPatients []Patient) {
	for _, patient := range patients {
		if patient.disease == disease {
			filteredPatients = append(filteredPatients, patient)
		}
	}
	return filteredPatients
}
