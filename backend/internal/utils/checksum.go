package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetUserHash(username string, password string, canteen string) string {
	userHashSHA := sha256.Sum256([]byte(username + password + canteen))
	userHashString := hex.EncodeToString(userHashSHA[:])

	return userHashString
}
