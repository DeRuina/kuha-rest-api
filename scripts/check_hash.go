package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	rawToken := "raw_token"
	expectedHash := "hashed_token"

	hash := sha256.Sum256([]byte(rawToken))
	computed := hex.EncodeToString(hash[:])

	fmt.Println("Computed hash:", computed)
	fmt.Println("Expected hash:", expectedHash)
	fmt.Println("Match?", computed == expectedHash)
}
