package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
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

		// Custom validation
		validate.RegisterValidation("key", func(fl validator.FieldLevel) bool {
			value := fl.Field().String()
			match, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, value)
			return match
		})
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

// Converts a string pointer to sql.NullString
func NullStringPtr(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

// Converts a float64 pointer to sql.NullFloat64
func NullFloat64Ptr(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

// Converts an int32 pointer to sql.NullInt32
func NullInt32Ptr(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

// Converts a Go type to a more user-friendly string representation
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

// ParseTimestamp parses a required RFC3339 timestamp string
func ParseTimestamp(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, ErrInvalidDate
	}
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, ErrInvalidTimeStamp
	}
	return t, nil
}

// ParseTimestampPtr parses a *string RFC3339 timestamp or returns nil
func ParseTimestampPtr(value *string) (*time.Time, error) {
	if value == nil || *value == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *value)
	if err != nil {
		return nil, ErrInvalidTimeStamp
	}
	return &t, nil
}

// ParseRawJSON converts a string pointer to pqtype.NullRawMessage
func ParseRawJSON(s *string) pqtype.NullRawMessage {
	if s == nil || *s == "" {
		return pqtype.NullRawMessage{}
	}
	return pqtype.NullRawMessage{Valid: true, RawMessage: json.RawMessage(*s)}
}

// ParseUUIDPtr parses a UUID string pointer to *uuid.NullUUID
func ParseUUIDPtr(s *string) (uuid.NullUUID, error) {
	if s == nil || *s == "" {
		return uuid.NullUUID{Valid: false}, nil
	}
	parsed, err := uuid.Parse(*s)
	if err != nil {
		return uuid.NullUUID{}, ErrInvalidUUID
	}
	return uuid.NullUUID{UUID: parsed, Valid: true}, nil
}

// Converts a *bool to sql.NullBool
func NullBoolPtr(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

// Converts a *time.Time to sql.NullTime
func NullTimePtr(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// ParseRequiredJSON converts a required string to json.RawMessage
func ParseRequiredJSON(s string) json.RawMessage {
	if s == "" {
		return json.RawMessage("null")
	}
	return json.RawMessage(s)
}

// Float64PtrOrNil converts sql.NullFloat64 to *float64
func Float64PtrOrNil(f sql.NullFloat64) *float64 {
	if f.Valid {
		return &f.Float64
	}
	return nil
}

// Int32PtrOrNil converts sql.NullInt32 to *int32
func Int32PtrOrNil(i sql.NullInt32) *int32 {
	if i.Valid {
		return &i.Int32
	}
	return nil
}

// FormatDatePtr formats sql.NullTime to a string pointer in "YYYY-MM-DD" format
func FormatDatePtr(t sql.NullTime) *string {
	if !t.Valid {
		return nil
	}
	str := t.Time.Format("2006-01-02")
	return &str
}

// StringPtrOrNil converts sql.NullString to *string
func StringPtrOrNil(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

// RawMessagePtrOrNil converts pqtype.NullRawMessage to *string
func RawMessagePtrOrNil(rm pqtype.NullRawMessage) *string {
	if rm.Valid && len(rm.RawMessage) > 0 {
		str := string(rm.RawMessage)
		return &str
	}
	return nil
}
