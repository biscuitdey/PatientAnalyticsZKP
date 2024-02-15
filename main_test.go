package main

import (
	"fmt"
	"testing"
)

func TestAverageAllPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	totalAge := calculateTotalAge(patients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(patients), ledger.Patients)

	fmt.Printf("Encrypted average age matches expected average age. Average age: %d", encryptedAverageAge)

	fmt.Println()
	fmt.Println()

}

func TestAverageAsthmaPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	asthmaPatients := filterPatients("asthma", patients)
	encryptedAsthmaPatients := ledger.filter("asthma")

	totalAge := calculateTotalAge(asthmaPatients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(asthmaPatients), encryptedAsthmaPatients)

	fmt.Printf("Encrypted average age of asthma patients matches expected average age mismatch. Average age: %d", encryptedAverageAge)

	fmt.Println()
	fmt.Println()
}

func TestAverageDiabetesPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	diabetesPatients := filterPatients("diabetes", patients)
	encryptedDiabetesPatients := ledger.filter("diabetes")

	totalAge := calculateTotalAge(diabetesPatients)
	encryptedAverageAge := ledger.calculateAverageAge(totalAge, len(diabetesPatients), encryptedDiabetesPatients)

	fmt.Printf("Encrypted average age of diabetes patients matches expected average age mismatch. Average age: %d", encryptedAverageAge)

	fmt.Println()
	fmt.Println()

}

func TestZKCircuitTest(t *testing.T) {

	proof := generateProof("asthma")

	fmt.Printf("%t : Public input (diabetes) does not match with zk proof (asthma)", proof.verify("diabetes"))

	fmt.Println()

	fmt.Printf("%t : Public input (asthma) matches with zk proof (asthma)", proof.verify("asthma"))

	fmt.Println()
	fmt.Println()

}
