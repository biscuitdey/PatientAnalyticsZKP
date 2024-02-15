package main

import (
	"crypto/sha256"
	"math/big"
)

func main() {

	//Patients
	patients := []Patient{
		Patient{
			disease: "asthma",
			age:     36,
		}, Patient{
			disease: "diabetes",
			age:     57,
		}, Patient{
			disease: "diabetes",
			age:     44,
		}, Patient{
			disease: "asthma",
			age:     59,
		}, Patient{
			disease: "asthma",
			age:     82,
		}, Patient{
			disease: "asthma",
			age:     37,
		}, Patient{
			disease: "diabetes",
			age:     53,
		}, Patient{
			disease: "asthma",
			age:     40,
		}, Patient{
			disease: "diabetes",
			age:     25,
		}, Patient{
			disease: "asthma",
			age:     76,
		}, Patient{
			disease: "asthma",
			age:     63,
		},
	}

	//TEST
	//1. Calculate average age of patients (unencrypted)

	//2. Calculate average age of patients (encrypted)
	//2.1. Create ledger
	//2.2. Create encryptedPatient
	//2.3. Add encryptedPatient to ledger
	//2.4. Get average get of patients with specific disease

	//3. Open the commitment to verify --> unencrypted average age matches encrypted average age

}

func hash(data string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte("string"))
	md := hasher.Sum(nil)

	return md
}

func hashToBigInt(data []byte) *big.Int {
	dataBigInt := new(big.Int).SetBytes(data)
	return dataBigInt
}
