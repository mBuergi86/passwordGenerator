package main

import (
	// "flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
	//	length := flag.Int("length", 8, "Password length")
	//	isDigit := flag.Bool("digit", true, "Include digits")
	//	isSpecial := flag.Bool("special", true, "Include special characters")
	//	isLowercase := flag.Bool("lowercase", true, "Include lowercase characters")
	//	isUppercase := flag.Bool("uppercase", true, "Include uppercase characters")

	//	flag.Parse()

	if len(os.Args) != 6 {
		fmt.Println("Usage: password-generator <length> <isDigit> <isSpecial> <isLowercase> <isUppercase>")
		fmt.Println("Example: password-generator 8 true true true true")
		os.Exit(1)
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid length")
		os.Exit(1)
	}

	isDigit, err := strconv.ParseBool(os.Args[2])
	if err != nil {
		fmt.Println("Invalid isDigit")
		os.Exit(1)
	}

	isSpecial, err := strconv.ParseBool(os.Args[3])
	if err != nil {
		fmt.Println("Invalid isSpecial")
		os.Exit(1)
	}

	isLowercase, err := strconv.ParseBool(os.Args[4])
	if err != nil {
		fmt.Println("Invalid isLowercase")
		os.Exit(1)
	}

	isUppercase, err := strconv.ParseBool(os.Args[5])
	if err != nil {
		fmt.Println("Invalid isUppercase")
		os.Exit(1)
	}

	// password := passwordGenerate(*length, *isDigit, *isSpecial, *isLowercase, *isUppercase)
	password := passwordGenerate(length, isDigit, isSpecial, isLowercase, isUppercase)

	fmt.Printf("Generated Password: %s\n", password)
}
