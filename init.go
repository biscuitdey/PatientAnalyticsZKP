package main

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/mit-dci/zksigma"
	"github.com/mit-dci/zksigma/btcec"
)

var PatientLedgerCurve zksigma.ZKPCurveParams

func generateH2tothe() []zksigma.ECPoint {
	Hslice := make([]zksigma.ECPoint, 64)
	for i := range Hslice {
		m := big.NewInt(1 << uint(i))
		Hslice[i].X, Hslice[i].Y = PatientLedgerCurve.C.ScalarBaseMult(m.Bytes())
	}
	return Hslice
}

func init() {
	s256 := sha256.New()

	curValue := btcec.S256().Gx
	s256.Write(new(big.Int).Add(curValue, big.NewInt(2)).Bytes())

	potentialXValue := make([]byte, 33)
	binary.LittleEndian.PutUint32(potentialXValue, 2)
	for i, elem := range s256.Sum(nil) {
		potentialXValue[i+1] = elem
	}

	H, err := btcec.ParsePubKey(potentialXValue, btcec.S256())
	if err != nil {
		panic(err)
	}
	PatientLedgerCurve = zksigma.ZKPCurveParams{
		C: btcec.S256(),
		G: zksigma.ECPoint{btcec.S256().Gx, btcec.S256().Gy},
		H: zksigma.ECPoint{H.X, H.Y},
	}
	PatientLedgerCurve.HPoints = generateH2tothe()
}
