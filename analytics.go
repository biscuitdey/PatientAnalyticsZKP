package main

import (
	"github.com/mit-dci/zksigma"
)

func computeSum(encryptedPatients []EncryptedPatient) (totalAge *AgeCommitment) {
	totalAgeCommitment := zksigma.Zero
	totalRandomToken := zksigma.Zero

	for i := 0; i < len(encryptedPatients); i++ {
		encryptedPatient := &encryptedPatients[i]
		totalAgeCommitment = PatientLedgerCurve.Add(totalAgeCommitment, encryptedPatient.age.commitment)
		totalRandomToken = PatientLedgerCurve.Add(totalRandomToken, encryptedPatient.age.randomValueToken)

	}
	return &AgeCommitment{
		commitment:       totalAgeCommitment,
		randomValueToken: totalRandomToken,
	}
}

func (ledger *Ledger) computeAverage(totalAge int, totalAgeCommitment AgeCommitment, totalPatients int) int {
	if verifyAgeCommitment(totalAge, &totalAgeCommitment, ledger.Key.pk, ledger.Key.sk) {
		return totalAge / totalPatients
	}

	return 0
}
