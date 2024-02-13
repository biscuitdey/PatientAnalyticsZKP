package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

func (ledger *Ledger) filter(disease string) (filteredPatients []EncryptedPatient) {

	for _, encryptedPatient := range ledger.Patients {
		//TODO: Check encryptedPatient struct name
		if encryptedPatient.diseaseZKProof.verify(disease) {
			filteredPatients = append(filteredPatients, encryptedPatient)
		}
	}
	return
}

func (ledger *Ledger) computeAverage(filteredPatients []EncryptedPatient) zksigma.ECPoint {
	comms := zksigma.Zero

	for i := 0; i < len(filteredPatients); i++ {
		encryptedPatient := &filteredPatients[i]
		comms = ZKLedgerCurve.Add(comms, encryptedPatient.Comm)

	}

	averageScalarInt := int64(1 / len(filteredPatients))
	averageScalarBigInt := new(big.Int).SetInt64(averageScalarInt)

	average := ZKLedgerCurve.Mult(comms, averageScalarBigInt)

	return average
}
