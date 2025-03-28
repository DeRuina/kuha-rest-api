package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// query timeout duration
const QueryTimeout = 7 * time.Second

// Validator to be initialized once
var validate *validator.Validate
var once sync.Once

// Validator getter
func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())
	})
	return validate
}

// Parse UUID from string
func ParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, ErrInvalidUUID
	}
	return parsedUUID, nil
}

// ParseDate converts a string (YYYY-MM-DD) to time.Time
func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}

	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, ErrInvalidDate
	}
	return parsedTime, nil
}

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

// Returns nil if the string is empty, otherwise returns the string pointer
func NilIfEmpty(s *string) *string {
	if s == nil || *s == "" {
		return nil
	}
	return s
}

// Converts an empty string to NULL for SQL compatibility
func NullTimeIfEmpty(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// Converts a string into sql.NullString with Valid=true
func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

// Checks if only allowed parameters are used in the request.
func ValidateParams(r *http.Request, allowedParams []string) error {
	allowed := make(map[string]bool)
	for _, param := range allowedParams {
		allowed[param] = true
	}

	var invalidParams []string
	for param := range r.URL.Query() {
		if !allowed[param] {
			invalidParams = append(invalidParams, param)
		}
	}

	if len(invalidParams) > 0 {
		return fmt.Errorf("invalid query parameters: %s. Allowed parameters are: %s",
			strings.Join(invalidParams, ", "),
			strings.Join(allowedParams, ", "))
	}

	return nil
}
