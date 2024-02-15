package main

import "testing"

func TestAverageAllPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	expectedAverageAge := calculateAverageAge(patients)
	encryptedAverageAge := calculateEncryptedAverageAge(ledger.Patients)
	isVerified := verifyAgeCommitment(expectedAverageAge, encryptedAverageAge)
	if !isVerified {
		t.Errorf("Encrypted average age doesn't match expected average age.")
	}

	t.Logf("Encrypted average age matches expected average age mismatch. Average age: %d", expectedAverageAge)

}

func TestAverageAsthmaPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	asthmaPatients := filterPatients("asthma", patients)
	encryptedAsthmaPatients := ledger.filter("asthma")

	expectedAverageAge := calculateAverageAge(asthmaPatients)
	encryptedAverageAge := calculateEncryptedAverageAge(encryptedAsthmaPatients)

	isVerified := verifyAgeCommitment(expectedAverageAge, encryptedAverageAge)
	if !isVerified {
		t.Errorf("Encrypted average age of asthma patients doesn't match expected average age.")
	}

	t.Logf("Encrypted average age of asthma patients matches expected average age mismatch. Average age: %d", expectedAverageAge)

}

func TestAverageDiabetesPatients(t *testing.T) {

	patients, ledger := SetupTestData()

	diabetesPatients := filterPatients("diabetes", patients)
	encryptedDiabetesPatients := ledger.filter("diabetes")

	expectedAverageAge := calculateAverageAge(diabetesPatients)
	encryptedAverageAge := calculateEncryptedAverageAge(encryptedDiabetesPatients)

	isVerified := verifyAgeCommitment(expectedAverageAge, encryptedAverageAge)
	if !isVerified {
		t.Errorf("Encrypted average age of diabetes patients doesn't match expected average age.")
	}

	t.Logf("Encrypted average age of diabetes patients matches expected average age mismatch. Average age: %d", expectedAverageAge)
}
