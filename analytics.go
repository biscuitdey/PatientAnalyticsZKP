package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

func computeAverage(encryptedPatients []EncryptedPatient) zksigma.ECPoint {
	totalAge := zksigma.Zero

	for i := 0; i < len(encryptedPatients); i++ {
		encryptedPatient := &encryptedPatients[i]
		totalAge = PatientLedgerCurve.Add(totalAge, encryptedPatient.age.commitment)

	}

	totalPatients := int64(1 / len(encryptedPatients)) // 1 / totalPatients
	totalPatientsBigInt := new(big.Int).SetInt64(totalPatients)

	average := PatientLedgerCurve.Mult(totalAge, totalPatientsBigInt)

	return average
}
