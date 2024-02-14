package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

func (ledger *Ledger) filter(disease string) (filteredPatients []EncryptedPatient) {

	for _, encryptedPatient := range ledger.Patients {
		//TODO: Check encryptedPatient struct name
		if encryptedPatient.disease.verify(disease) {
			filteredPatients = append(filteredPatients, encryptedPatient)
		}
	}
	return
}

func (ledger *Ledger) computeAverage(filteredPatients []EncryptedPatient) zksigma.ECPoint {
	totalAge := zksigma.Zero

	for i := 0; i < len(filteredPatients); i++ {
		encryptedPatient := &filteredPatients[i]
		totalAge = ZKLedgerCurve.Add(totalAge, encryptedPatient.age)

	}

	totalPatients := int64(1 / len(filteredPatients)) // 1 / totalPatients
	totalPatientsBigInt := new(big.Int).SetInt64(totalPatients)

	average := ZKLedgerCurve.Mult(totalAge, totalPatientsBigInt)

	return average
}
