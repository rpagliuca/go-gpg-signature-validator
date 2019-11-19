package main

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	message := `-----BEGIN PGP MESSAGE-----

owEBYgKd/ZANAwAIAXf0GgpR5Yx1AcsbYgBd092wV2hvIGVsc2UgYnV0IG15c2Vs
Zj8KiQIzBAABCAAdFiEEyUqsAqTRGlH8yS3Nd/QaClHljHUFAl3T3bAACgkQd/Qa
ClHljHU5ig//RtXMmtSGA3+ZYrTOEEr0gK4pOxvla1ZvWHlHZQqYV3uRWxm8c/mb
M6ub5pBD/2dosmdbU8bFvgpPQ3xPN6wr/5J40sZKvFSgxa33LVBoVXcgTC8jeDd/
jxQ2dMagPMoX3Q+8B9OI0KDFiIgxe9+rAcL259l/p6DZGx5QF+R+eQySUZ1UjUcn
1L7A6nAUcXAeAPFPnAHjbTrc35fsR2ejZs4WJvUwAbcWFFPB2djGBDiEhOPewcYP
RmsXwWa+GIwf97IxcNvdDza4b9fsxw0MHCzteTzrMmAzvsBbhEnqdAg71n97Pesi
liKWTD8IAJcRLrigdayRCKquf1yxdfeEAsSaS+GcCaCTdtu/2lNPABaAW8bUU6UF
qZZWLB7NXOMNV0gZCo08A65s1DVIXVbDRQtBytCGt/cJFiIQlXr6dICsZihShWdO
2nER2NIzeD+zLcHl9t/aPQ/+xWynat9rYaftaSkD+/ZFc/irsLprG76agxpD74Af
sUvtsd5mee7+kOAWxfSE2xfNHtmoROZJ8uEtdKqXtq4Ub5Ko8nvgjdYkEbtqdI2e
rz+N4kNZZz3kQsbGDMt1bKZqc8b2zJ/CyJ/mHyI5IIx7wEEG4FnvSVn4d7qsEShh
2bv1ekbtDudCFmTQ/DfHhGki2tMzbcLvPtAhJaIeggB1YeLUGatLCu4=
=71ou
-----END PGP MESSAGE-----`
	isValid, contents, err := GpgSignatureIsValid(message, "C94AAC02A4D11A51FCC92DCD77F41A0A51E58C75")
	fmt.Println(contents)
	if !isValid {
		t.Error(err)
	}
	if contents != "Who else but myself?\n" {
		t.Error("Invalid message contents")
	}
}
