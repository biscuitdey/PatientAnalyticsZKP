package main

import (
	"math/big"

	"github.com/mit-dci/zksigma"
)

type Ledger struct {
	Patients []EncryptedPatient
	Key      *Key
}

type Key struct {
	pk zksigma.ECPoint
	sk *big.Int
}

func generateKey() *Key {
	newPk, newSK := zksigma.KeyGen(PatientLedgerCurve.C, PatientLedgerCurve.H)

	key := &Key{
		pk: newPk,
		sk: newSK,
	}

	return key
}

func MakeLedger() *Ledger {
	ledger := &Ledger{
		Patients: make([]EncryptedPatient, 0),
		Key:      generateKey(),
	}
	return ledger
}

func (ledger *Ledger) add(encryptedPatient *EncryptedPatient) int {
	encryptedPatient.id = len(ledger.Patients)
	ledger.Patients = append(ledger.Patients, *encryptedPatient)
	return encryptedPatient.id
}

func (ledger *Ledger) filter(disease string) (filteredPatients []EncryptedPatient) {

	for _, encryptedPatient := range ledger.Patients {
		if encryptedPatient.disease.verify(disease) {
			filteredPatients = append(filteredPatients, encryptedPatient)
		}
	}
	return
}
