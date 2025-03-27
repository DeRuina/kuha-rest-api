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
	"github.com/sqlc-dev/pqtype"
)

// Map of clients to their assigned roles
// {clientName: {roles}}
var clientsToSeed = map[string][]string{}

// Seed inserts predefined clients with hashed client_tokens into the database
func Seed(store store.Auth, db *sql.DB) {
	ctx := context.Background()
	queries := store.Queries()

	// Delete all existing clients and reset their ID sequence
	if _, err := db.ExecContext(ctx, `DELETE FROM clients`); err != nil {
		log.Fatalf("Failed to delete clients: %v", err)
	}
	if _, err := db.ExecContext(ctx, `ALTER SEQUENCE clients_id_seq RESTART WITH 1`); err != nil {
		log.Fatalf("Failed to reset clients_id_seq: %v", err)
	}
	log.Println("Clients table cleared and sequence reset.")

	// Delete all token logs and reset their ID sequence
	if _, err := db.ExecContext(ctx, `DELETE FROM token_logs`); err != nil {
		log.Fatalf("Failed to delete token logs: %v", err)
	}
	if _, err := db.ExecContext(ctx, `ALTER SEQUENCE token_logs_id_seq RESTART WITH 1`); err != nil {
		log.Fatalf("Failed to reset token_logs_id_seq: %v", err)
	}
	log.Println("Token logs cleared and sequence reset.")

	// Create or overwrite the output file
	file, err := os.Create("client_tokens.txt")
	if err != nil {
		log.Fatalf("Failed to create client_tokens.txt: %v", err)
	}
	defer file.Close()

	for clientName, roles := range clientsToSeed {
		// Generate token pair
		rawToken, hashedToken, err := generateSecureTokenPair(32)
		if err != nil {
			log.Printf("Failed to generate token for %s: %v", clientName, err)
			continue
		}

		// Insert client
		err = queries.CreateClient(ctx, authsqlc.CreateClientParams{
			ClientName:  clientName,
			ClientToken: hashedToken,
			Role:        roles,
		})
		if err != nil {
			log.Printf("Error inserting client %s: %v", clientName, err)
			continue
		}

		// Log issued client_token
		metadata := pqtype.NullRawMessage{Valid: true}
		if err := metadata.Scan([]byte(`"seeding"`)); err != nil {
			log.Printf("Failed to prepare metadata for %s: %v", clientName, err)
		}

		err = queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
			ClientToken: hashedToken,
			TokenType:   "client",
			Action:      "issued",
			Token:       sql.NullString{String: hashedToken, Valid: true},
			IpAddress:   sql.NullString{},
			UserAgent:   sql.NullString{},
			Metadata:    metadata,
		})
		if err != nil {
			log.Printf("Failed to log client_token for %s: %v", clientName, err)
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

	raw := hex.EncodeToString(bytes)      // Original token
	hash := sha256.Sum256([]byte(raw))    // SHA-256 hash
	hashed := hex.EncodeToString(hash[:]) // Convert hash to hex

	return raw, hashed, nil
}
