package logic

import (
	"math/rand"
	"strconv"
	"strings"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "1234567890"
	specialChars     = "!#$&*_"
)

func GeneratedPassword(lenPassword string) string {
	lenInt, _ := strconv.Atoi(lenPassword)  // string to number

	allChars := lowercaseLetters + uppercaseLetters + digits + specialChars

	for {
		password := make([]byte, lenInt)
		for i := range password {
			password[i] = allChars[rand.Intn(len(allChars))]
		}
		passwordStr := string(password)
		if IsValidPassword(passwordStr) {
			return passwordStr
		}
	}
}

func IsValidPassword(password string) bool {
	hasLowercase := false
	hasUppercase := false
	hasDigit := false
	hasSpecialChars := false

	for _, char := range password {
		if strings.Contains(lowercaseLetters, string(char)) {
			hasLowercase = true
		} else if strings.Contains(uppercaseLetters, string(char)) {
			hasUppercase = true
		} else if strings.Contains(digits, string(char)) {
			hasDigit = true
		} else if strings.Contains(specialChars, string(char)) {
			hasSpecialChars = true
		}
	}
	return hasLowercase && hasUppercase && hasDigit && hasSpecialChars
}
