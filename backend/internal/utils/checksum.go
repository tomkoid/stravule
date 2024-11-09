package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GetUserHash(username string, password string, canteen string) string {
	userHashSHA := sha256.Sum256([]byte(strings.TrimSpace(username) + strings.TrimSpace(password) + strings.TrimSpace(canteen)))
	userHashString := hex.EncodeToString(userHashSHA[:])

	return userHashString
}
