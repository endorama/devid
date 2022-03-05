package utils

import petname "github.com/dustinkirkland/golang-petname"

const passphraseLength = 6

func GeneratePassphrase() string {
	petname.NonDeterministicMode()

	return petname.Generate(passphraseLength, "-")
}
