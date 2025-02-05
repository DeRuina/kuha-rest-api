package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Parse UUID from string
func ParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format")
	}
	return parsedUUID, nil
}

// Parse date string (YYYY-MM-DD) to time.Time
func ParseDate(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return parsedTime, nil
}

// Parse date string (YYYY-MM-DD) to Unix timestamp (float64)
func ParseTimestamp(dateStr string) (float64, error) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, errors.New("invalid timestamp format, expected YYYY-MM-DD")
	}
	return float64(parsedTime.Unix()), nil
}

// Convert string to JSON raw message
func ParseJSONKey(key string) json.RawMessage {
	return json.RawMessage(`"` + key + `"`)
}
