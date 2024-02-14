package main

import (
	"crypto/sha256"
	"math/big"
)

func main() {

	//TEST
	//1. Calculate average age of patients (unencrypted)

	//2. Calculate average age of patients (encrypted)
	//2.1. Create ledger
	//2.2. Create encryptedPatient
	//2.3. Add encryptedPatient to ledger
	//2.4. Get average get of patients with specific disease

	//3. Open the commitment to verify --> unencrypted average age matches encrypted average age
	// zksigma.Open(a, b, c, average)

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
