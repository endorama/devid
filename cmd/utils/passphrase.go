package utils

import petname "github.com/dustinkirkland/golang-petname"

func GeneratePassphrase() string {
	petname.NonDeterministicMode()
	passphraseLength := 6
	return petname.Generate(passphraseLength, "-")
}
