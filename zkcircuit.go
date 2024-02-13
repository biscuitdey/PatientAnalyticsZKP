package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type ZeroKnowledge struct {
	proof        groth16.Proof
	verifyingKey groth16.VerifyingKey
}

// PatientDisease == InputDisease
type DiseaseCircuit struct {
	PatientDisease frontend.Variable `gnark:"x"`
	InputDisease   frontend.Variable `gnark:",public"`
}

func (circuit *DiseaseCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(circuit.PatientDisease, circuit.InputDisease)
	return nil
}

func (zk *ZeroKnowledge) generateProof(disease string) *ZeroKnowledge {
	// compiles our circuit into a R1CS
	var circuit DiseaseCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	// groth16 zkSNARK: Setup
	pk, vk, _ := groth16.Setup(ccs)

	//Disease name - hash and convert to bigint
	diseaseHash := hash(disease)
	diseaseBigInt := hashToBigInt(diseaseHash)

	// witness definition
	assignment := DiseaseCircuit{PatientDisease: diseaseBigInt, InputDisease: diseaseBigInt}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	// groth16: Prove & Verify
	proof, _ := groth16.Prove(ccs, pk, witness)

	zkValues := &ZeroKnowledge{
		proof:        proof,
		verifyingKey: vk,
	}

	return zkValues
}

func (zk *ZeroKnowledge) verify(disease string) bool {
	//Disease name - hash and convert to bigint
	diseaseHash := hash(disease)
	diseaseBigInt := hashToBigInt(diseaseHash)

	//public witness definition
	assignment := DiseaseCircuit{InputDisease: diseaseBigInt}
	publicWitness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField(), frontend.PublicOnly())

	err := groth16.Verify(zk.proof, zk.verifyingKey, publicWitness)

	if err != nil {
		return false
	}
	return true
}