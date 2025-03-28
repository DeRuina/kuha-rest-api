package authn

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

// Load the JWT env
var (
	jwtSecret   = []byte(env.GetString("JWT_SECRET", ""))
	jwtIssuer   = env.GetString("JWT_ISSUER", "")
	jwtAudience = env.GetString("JWT_AUDIENCE", "")
)

// GenerateRandomToken returns a secure 32-byte (256-bit) random token as hex string
func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateJWT creates a signed JWT with roles and specified expiry duration
func GenerateJWT(clientName string, roles []string, duration time.Duration) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub":   clientName,
		"roles": roles,
		"exp":   now.Add(duration).Unix(),
		"iat":   now.Unix(),
		"iss":   jwtIssuer,
		"aud":   jwtAudience,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
