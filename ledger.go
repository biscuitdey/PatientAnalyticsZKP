package main

type Ledger struct {
	Patients []EncryptedPatient
}

func MakeLedger() *Ledger {
	ledger := &Ledger{
		Patients: make([]EncryptedPatient, 0),
	}
	return ledger
}

func (ledger *Ledger) add(ecryptedPatient *EncryptedPatient) int {
	ecryptedPatient.Index = len(ledger.Patients)
	ledger.Patients = append(ledger.Patients, *ecryptedPatient)
	return ecryptedPatient.Index
}
