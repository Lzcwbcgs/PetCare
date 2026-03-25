package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

func verifyPassword(plainPassword string, storedCandidates ...string) bool {
	hashedPassword := hashPassword(plainPassword)
	for _, candidate := range storedCandidates {
		if candidate == "" {
			continue
		}
		if candidate == plainPassword || candidate == hashedPassword {
			return true
		}
	}
	return false
}

