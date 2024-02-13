package main

import (
	"crypto/sha256"
	"math/big"

	"github.com/consensys/gnark/frontend"
)

type CubicCircuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints
// x**3 + x + 5 == y
func (circuit *CubicCircuit) Define(api frontend.API) error {
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)
	api.AssertIsEqual(circuit.Y, api.Add(x3, circuit.X, 5))
	return nil
}

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
