package utils

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Parse UUID from string
func ParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, ErrInvalidUUID
	}
	return parsedUUID, nil
}

// // ParseDate converts a string (YYYY-MM-DD) to time.Time
// func ParseDate(dateStr string) (time.Time, error) {
// 	if dateStr == "" {
// 		return time.Time{}, errors.New("date string is empty")
// 	}

// 	parsedTime, err := time.Parse("2006-01-02", dateStr)
// 	if err != nil {
// 		return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
// 	}
// 	return parsedTime, nil
// }

// ParseDatePtr converts a string (YYYY-MM-DD) to *time.Time
func ParseDatePtr(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}
	return &parsedTime, nil
}

// // Convert string to JSON raw message
// func ParseJSONKey(key string) json.RawMessage {
// 	return json.RawMessage(`"` + key + `"`)
// }

// // NilIfEmpty returns nil if the string is empty, otherwise returns the string pointer
// func NilIfEmpty(s *string) *string {
// 	if s == nil || *s == "" {
// 		return nil
// 	}
// 	return s
// }

// NullTimeIfEmpty converts an empty string to NULL for SQL compatibility
func NullTimeIfEmpty(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}
