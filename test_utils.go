package main

import (
	"github.com/mit-dci/zksigma"
)

func SetupTestData() ([]Patient, *Ledger) {
	//Make ledger
	ledger := MakeLedger()

	//Patient data
	patients := []Patient{
		{
			disease: "asthma",
			age:     36,
		}, {
			disease: "diabetes",
			age:     57,
		}, {
			disease: "diabetes",
			age:     44,
		}, {
			disease: "asthma",
			age:     59,
		}, {
			disease: "asthma",
			age:     82,
		}, {
			disease: "asthma",
			age:     37,
		}, {
			disease: "diabetes",
			age:     53,
		}, {
			disease: "asthma",
			age:     40,
		}, {
			disease: "diabetes",
			age:     25,
		}, {
			disease: "asthma",
			age:     76,
		}, {
			disease: "asthma",
			age:     63,
		},
	}

	//Encrypt patient data and add to ledger
	for _, patient := range patients {
		encryptedPatient := MakeEncryptedPatient(patient.disease, patient.age)
		ledger.add(encryptedPatient)
	}

	return patients, ledger
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
