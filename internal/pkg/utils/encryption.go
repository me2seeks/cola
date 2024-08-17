package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Encrypt encrypts the password with the salt.
func Encrypt(password, salt string) string {
	combined := password + salt

	hash := sha256.New()
	hash.Write([]byte(combined))

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
