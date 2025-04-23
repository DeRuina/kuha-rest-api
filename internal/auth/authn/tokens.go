package authn

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Load the JWT env
type JWTConfig struct {
	Secret   []byte
	Issuer   string
	Audience string
}

var (
	jwtSecret   []byte
	jwtIssuer   string
	jwtAudience string
)

func LoadJWTConfig(cfg JWTConfig) {
	jwtSecret = cfg.Secret
	jwtIssuer = cfg.Issuer
	jwtAudience = cfg.Audience
}

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

func ValidateJWT(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	},
		jwt.WithAudience(jwtAudience),
		jwt.WithIssuer(jwtIssuer),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, claims, nil
}
