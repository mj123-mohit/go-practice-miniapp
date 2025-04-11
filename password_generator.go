package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generatePassword(length int) string {
	// Define character sets for the password
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	allChars := lower + upper + digits + special

	// Initialize an empty password
	var password string

	for i := 0; i < length; i++ {
		// Generate a random index within the valid character range
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		password += string(allChars[randomIndex.Int64()])
	}

	return password
}

func main() {
	password := generatePassword(12) // Generate a 12-character password
	fmt.Println("Generated Password: ", password)
}
