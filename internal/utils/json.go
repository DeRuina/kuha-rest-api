package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func goTypeToFriendlyType(goType string) string {
	switch goType {
	case "string":
		return "a string"
	case "int", "int32", "int64", "float32", "float64":
		return "a number"
	case "Time":
		return "a date (YYYY-MM-DD)"
	default:
		return "the correct type"
	}
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	if r.Header.Get("X-Was-Gzipped") != "true" {
		maxBytes := int64(50 * 1024 * 1024) // 50MB for regular content only
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(data); err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			return &InvalidFieldTypeError{
				Field:        e.Field,
				ExpectedType: goTypeToFriendlyType(e.Type.Name()),
				ActualType:   e.Value,
			}

		case *json.InvalidUnmarshalError:
			return ErrInvalidValue
		default:
			if errors.Is(err, io.EOF) {
				return fmt.Errorf("request body is required")
			}
			return err
		}
	}

	return nil
}

func WriteJSONError(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	switch msg := message.(type) {
	case string:
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]string{
				{"error": msg},
			},
		})
	case map[string]string:
		var errorList []map[string]string
		for key, val := range msg {
			errorList = append(errorList, map[string]string{key: val})
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"errors": errorList})
	default:
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": []map[string]string{
				{"error": "An unexpected error occurred"},
			},
		})
	}
}
