package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func passwordGenerate(length int, isDigit bool, isSpecial bool, isLowercase bool, isUppercase bool) string {
	charset := ""

	if isDigit {
		charset += "0123456789"
	}

	if isSpecial {
		charset += `!@#$%^&*()_+{}|:<>?-=[]\;',./`
	}

	if isLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}

	if isUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	passwordBuilder := strings.Builder{}
	passwordBuilder.Grow(length)

	for i := 0; i < length; i++ {
		passwordBuilder.WriteByte(charset[rand.Intn(len(charset))])
	}

	return passwordBuilder.String()
}

func main() {
	length := 12
	fmt.Print(passwordGenerate(length, true, true, true, true))
}
