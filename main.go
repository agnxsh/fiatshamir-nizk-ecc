package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/util/random"
)

var rng = random.New()

func main () {
	suite := edwards25519.NewBlakeSHA256Ed25519()

	m:="Hello"

	argCount := len(os.Args[1:])

	if (argCount > 0) {
		m = string(os.Args[1])
	}

	message := [] byte(m)
	scalar_ := sha256.Sum256(message[:])

	x := suite.Scalar().SetBytes(scalar_[:32])

	G := suite.Point().Pick(rng)
	H := suite.Point().Pick(rng)

	fmt.Printf("Bob and Alice agree:\n G:\t%s\n H:\t%s\n\n", G, H)

	fmt.Printf("Bob's Password: \t%s\n", m)

	fmt.Printf("Bob's Secret (x): \t%s\n\n", x)

	xG := suite.Point().Mul(x,G)
	xH := suite.Point().Mul(x,H)

	fmt.Printf("Bob sends these values to Alice: \n xG: \t%s\n xH: \t%s\n\n", xG, xH)

	v := suite.Scalar().Pick(suite.RandomStream())
	vG := suite.Point().Mul(v,G)
	vH := suite.Point().Mul(v,H)


	h:= suite.Hash()
	xG.MarshalTo(h)
	xH.MarshalTo(h)
	vG.MarshalTo(h)g
	vH.MarshalTo(h)

	cb := h.Sum(nil)
	c := suite.Scalar().Pick(suite.XOF(cb))

	//Response from Alice
	r := suite.Scalar()
	r.Mul(x, c).Sub(v, r)



}
