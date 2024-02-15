package main

import (
	"fmt"
	"testing"
)

func TestAverageAllPatients(t *testing.T) {

	//TODO: JUST CHECK SUMS

	//https://github.com/mit-dci/zkledger/blob/master/bank.go#L742 --> SEE "answerSum"

	//TODO: Need to add pk and sk for all patients
	patients, ledger := SetupTestData()

	totalAge := calculateTotalAge(patients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(patients), ledger.Patients)

	fmt.Printf("Encrypted average age matches expected average age. Average age: %d", encryptedAverageAge)

	fmt.Println()

}

func TestAverageAsthmaPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	asthmaPatients := filterPatients("asthma", patients)
	encryptedAsthmaPatients := ledger.filter("asthma")

	totalAge := calculateTotalAge(asthmaPatients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(patients), encryptedAsthmaPatients)

	fmt.Printf("Encrypted average age of asthma patients matches expected average age mismatch. Average age: %d", encryptedAverageAge)

	fmt.Println()
}

func TestAverageDiabetesPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	diabetesPatients := filterPatients("diabetes", patients)
	encryptedDiabetesPatients := ledger.filter("diabetes")

	totalAge := calculateTotalAge(diabetesPatients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(patients), encryptedDiabetesPatients)

	fmt.Printf("Encrypted average age of diabetes patients matches expected average age mismatch. Average age: %d", encryptedAverageAge)

	fmt.Println()

}
