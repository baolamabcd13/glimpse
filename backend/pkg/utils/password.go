package utils

import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "fmt"
    "strings"

    "golang.org/x/crypto/argon2"
)

// PasswordConfig defines the parameters for Argon2id
type PasswordConfig struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

// DefaultPasswordConfig returns the default password hashing configuration
func DefaultPasswordConfig() *PasswordConfig {
    return &PasswordConfig{
        Memory:      64 * 1024,
        Iterations:  3,
        Parallelism: 2,
        SaltLength:  16,
        KeyLength:   32,
    }
}

// HashPassword generates a hash of the password using Argon2id
func HashPassword(password string) (string, error) {
    config := DefaultPasswordConfig()

    // Generate a random salt
    salt := make([]byte, config.SaltLength)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }

    // Hash the password
    hash := argon2.IDKey(
        []byte(password),
        salt,
        config.Iterations,
        config.Memory,
        config.Parallelism,
        config.KeyLength,
    )

    // Encode as base64
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)

    // Format the hash
    encodedHash := fmt.Sprintf(
        "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version, config.Memory, config.Iterations, config.Parallelism,
        b64Salt, b64Hash,
    )

    return encodedHash, nil
}

// VerifyPassword checks if the provided password matches the hash
func VerifyPassword(password, encodedHash string) (bool, error) {
    // Extract the parameters, salt and hash from the encoded hash
    parts := strings.Split(encodedHash, "$")
    if len(parts) != 6 {
        return false, fmt.Errorf("invalid hash format")
    }

    var version int
    _, err := fmt.Sscanf(parts[2], "v=%d", &version)
    if err != nil {
        return false, err
    }

    var memory, iterations uint32
    var parallelism uint8
    _, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
    if err != nil {
        return false, err
    }

    salt, err := base64.RawStdEncoding.DecodeString(parts[4])
    if err != nil {
        return false, err
    }

    hash, err := base64.RawStdEncoding.DecodeString(parts[5])
    if err != nil {
        return false, err
    }

    // Compute the hash of the provided password
    keyLength := uint32(len(hash))
    comparisonHash := argon2.IDKey(
        []byte(password),
        salt,
        iterations,
        memory,
        parallelism,
        keyLength,
    )

    // Compare the computed hash with the stored hash
    return subtle.ConstantTimeCompare(hash, comparisonHash) == 1, nil
}