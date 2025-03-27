package seed

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
)

// List of clients to seed into the auth database
var clientsToSeed = []string{}

// Seed inserts predefined clients with hashed client_tokens into the database,
func Seed(store store.Auth, db *sql.DB) {
	ctx := context.Background()
	queries := store.Queries()

	// Create or overwrite the output file
	file, err := os.Create("client_tokens.txt")
	if err != nil {
		log.Fatalf("Failed to create client_tokens.txt: %v", err)
	}
	defer file.Close()

	for _, clientName := range clientsToSeed {
		// Generate raw and hashed token
		rawToken, hashedToken, err := generateSecureTokenPair(32)
		if err != nil {
			log.Printf("Failed to generate token for %s: %v", clientName, err)
			continue
		}

		// Insert into DB
		err = queries.CreateClient(ctx, authsqlc.CreateClientParams{
			ClientName:  clientName,
			ClientToken: hashedToken,
			Role:        []string{},
		})

		if err != nil {
			log.Printf("Error inserting client %s: %v", clientName, err)
			continue
		}

		// Save raw token to file
		line := fmt.Sprintf("Client: %-15s Token: %s\n", clientName, rawToken)
		if _, err := file.WriteString(line); err != nil {
			log.Printf("Failed to write token to file for %s: %v", clientName, err)
		}

		fmt.Print(line)
	}

	log.Println("Seeding complete. Tokens written to client_tokens.txt")
}

// Create a random token and its SHA-256 hash.
func generateSecureTokenPair(length int) (rawToken string, hashedToken string, err error) {
	bytes := make([]byte, length)
	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", err
	}

	raw := hex.EncodeToString(bytes)      // 64 hex chars (32 bytes)
	hash := sha256.Sum256([]byte(raw))    // Hash the token
	hashed := hex.EncodeToString(hash[:]) // Convert hash to hex string

	return raw, hashed, nil
}
