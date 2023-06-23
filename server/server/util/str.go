package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	// Calculate the number of bytes needed for the random string
	numBytes := (length * 6) / 8

	// Generate random bytes
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a base64-encoded string
	randomString := base64.URLEncoding.EncodeToString(bytes)

	// Trim any padding characters from the end of the string
	randomString = randomString[:length]

	return randomString, nil
}
