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

func (ledger *Ledger) add(encryptedPatient *EncryptedPatient) int {
	encryptedPatient.id = len(ledger.Patients)
	ledger.Patients = append(ledger.Patients, *encryptedPatient)
	return encryptedPatient.id
}
