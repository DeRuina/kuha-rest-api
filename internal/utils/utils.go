package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

// query timeout duration
const QueryTimeout = 30 * time.Second

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

// ParseTimestampPtr parses a pointer to a timestamp string in various formats
func ParseTimestampPtr(value *string) (*time.Time, error) {
	if value == nil {
		return nil, nil
	}
	s := strings.TrimSpace(*value)
	if s == "" {
		return nil, nil
	}

	// allow lowercase 'z'
	if strings.HasSuffix(s, "z") {
		s = s[:len(s)-1] + "Z"
	}

	// RFC3339 (with or without fractional seconds)
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		u := t.UTC()
		return &u, nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		u := t.UTC()
		return &u, nil
	}

	// No timezone -> assume UTC
	// (T or space separator; with/without fractional seconds)
	for _, layout := range []string{
		"2006-01-02T15:04:05.999999999",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05.999999999",
		"2006-01-02 15:04:05",
	} {
		if t, err := time.ParseInLocation(layout, s, time.UTC); err == nil {
			u := t.UTC()
			return &u, nil
		}
	}

	return nil, ErrInvalidTimeStamp
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

// UUIDPtrToStringPtr converts uuid.NullUUID to *string
func UUIDPtrToStringPtr(u uuid.NullUUID) *string {
	if u.Valid {
		str := u.UUID.String()
		return &str
	}
	return nil
}

// BoolPtrOrNil converts sql.NullBool to *bool
func BoolPtrOrNil(b sql.NullBool) *bool {
	if b.Valid {
		return &b.Bool
	}
	return nil
}

// RawMessageToString converts json.RawMessage to string (for non-nullable fields)
func RawMessageToString(rm json.RawMessage) string {
	return string(rm)
}

// Int64PtrOrNil converts sql.NullInt64 to *int64
func Int64PtrOrNil(i sql.NullInt64) *int64 {
	if i.Valid {
		return &i.Int64
	}
	return nil
}

// Int16AsInt32PtrOrNil converts sql.NullInt16 to *int32 (friendlier JSON type)
func Int16AsInt32PtrOrNil(i sql.NullInt16) *int32 {
	if i.Valid {
		v := int32(i.Int16)
		return &v
	}
	return nil
}

// FormatTimestampPtr formats sql.NullTime to *string (RFC3339)
func FormatTimestampPtr(t sql.NullTime) *string {
	if !t.Valid {
		return nil
	}
	s := t.Time.Format(time.RFC3339)
	return &s
}

// ParsePositiveInt32 parses a decimal string into a positive int32.
func ParsePositiveInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil || v <= 0 {
		return 0, ErrInvalidParameter
	}
	return int32(v), nil
}

// ParseNonNegativeInt32 parses a decimal string into a non-negative int32.
func ParseNonNegativeInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil || v < 0 {
		return 0, ErrInvalidParameter
	}
	return int32(v), nil
}

// ParseSporttiID trims and validates a numeric sportti_id.
func ParseSporttiID(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", ErrInvalidParameter
	}
	if _, err := strconv.ParseUint(s, 10, 64); err != nil {
		return "", ErrInvalidParameter
	}
	return s, nil
}

// Converts an int32 value to sql.NullInt32 with Valid=true
func NullInt32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: true}
}

// Convert to *float64 for friendlier JSON.
func NullNumericToFloatPtr(ns sql.NullString) *float64 {
	if !ns.Valid || ns.String == "" {
		return nil
	}
	f, err := strconv.ParseFloat(ns.String, 64)
	if err != nil {
		return nil
	}
	return &f
}

// Converts *int64 to sql.NullInt64
func NullInt64Ptr(v *int64) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *v, Valid: true}
}

// Converts *int32 to sql.NullInt16
func NullInt16FromInt32Ptr(v *int32) sql.NullInt16 {
	if v == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{Int16: int16(*v), Valid: true}
}

// Safe deref helpers (defaulting to zero values)
func DerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
func DerefInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// Converts *float64 to sql.NullString for NUMERIC columns
// (complements NullNumericToFloatPtr)
func NullNumericFromFloat64Ptr(f *float64) sql.NullString {
	if f == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: strconv.FormatFloat(*f, 'f', -1, 64),
		Valid:  true,
	}
}

// More flexible timestamp parser: RFC3339, "YYYY-MM-DD HH:MM:SS UTC", or no-zone (UTC)
func ParseTimestampPtrFlexible(value *string) (*time.Time, error) {
	if value == nil || *value == "" {
		return nil, nil
	}
	if t, err := time.Parse(time.RFC3339, *value); err == nil {
		return &t, nil
	}
	if t, err := time.Parse("2006-01-02 15:04:05 MST", *value); err == nil {
		return &t, nil
	}
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", *value, time.UTC); err == nil {
		return &t, nil
	}
	return nil, ErrInvalidTimeStamp
}

// TimePtrOrNil converts sql.NullTime to *time.Time
func TimePtrOrNil(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func ParseRFC3339MinuteOrSecond(s string) (time.Time, error) {
	// try full RFC3339 first (with seconds)
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t.Truncate(time.Minute).UTC(), nil
	}
	// then try without seconds
	const noSec = "2006-01-02T15:04Z07:00"
	if t, err := time.Parse(noSec, s); err == nil {
		return t.Truncate(time.Minute).UTC(), nil
	}
	return time.Time{}, ErrInvalidTimeStamp
}

func ParsePositiveInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil || v <= 0 {
		return 0, ErrInvalidParameter
	}
	return v, nil
}
