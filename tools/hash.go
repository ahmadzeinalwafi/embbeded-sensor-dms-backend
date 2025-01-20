package tools

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// GenerateSalt generates a random salt in 16-byte
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword hashes a password using the Argon2id algorithm and returns a string representation of the hash
// and salt combined. The resulting format is "salt$hash", where both components are base64-encoded.
//
// Parameters:
//   - password: The plaintext password to hash.
//
// Returns:
//   - string: A base64-encoded string containing the salt and hash separated by "$".
//   - error: An error if the salt generation or hashing process fails.
//
// Argon2 Parameters:
//   - timeCost: Number of iterations (set to 1 for simplicity).
//   - memCost: Amount of memory (64 MB).
//   - parallelism: Number of threads (set to 4).
//   - keyLength: Length of the resulting hash (32 bytes).
func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	timeCost := uint32(1)
	memCost := uint32(64 * 1024)
	parallelism := uint8(4)
	keyLength := uint32(32)

	hash := argon2.IDKey([]byte(password), salt, timeCost, memCost, parallelism, keyLength)

	hashWithSalt := fmt.Sprintf("%s$%s", base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hash))
	return hashWithSalt, nil
}

// VerifyPassword verifies whether a provided plaintext password matches a previously hashed password.
// It rehashes the plaintext password using the same salt and Argon2 parameters and compares the resulting hash.
//
// Parameters:
//   - password: The plaintext password to verify.
//   - hashedPassword: The hashed password string in the format "salt$hash" where both are base64-encoded.
//
// Returns:
//   - bool: `true` if the password matches the hashed password, `false` otherwise.
//   - error: An error if the hashed password format is invalid or decoding fails.
func VerifyPassword(password, hashedPassword string) (bool, error) {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 2 {
		return false, errors.New("invalid hashed password format")
	}

	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, errors.New("invalid salt encoding")
	}

	storedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, errors.New("invalid hash encoding")
	}

	timeCost := uint32(1)
	memCost := uint32(64 * 1024)
	parallelism := uint8(4)
	keyLength := uint32(32)

	newHash := argon2.IDKey([]byte(password), salt, timeCost, memCost, parallelism, keyLength)

	if string(newHash) == string(storedHash) {
		return true, nil
	}
	return false, nil
}
