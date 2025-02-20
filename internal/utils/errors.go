package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Errors
var (
	ErrInvalidUUID      = errors.New("invalid UUID")
	ErrInvalidDate      = errors.New("invalid date: ensure the format is YYYY-MM-DD and values are realistic")
	ErrMissingUserID    = errors.New("user_id is required")
	ErrMissingDate      = errors.New("date is required")
	ErrInvalidParameter = errors.New("invalid parameter provided")
	ErrInvalidDateRange = errors.New("invalid date range")
	ErrInvalidChoice    = errors.New("invalid choice: must be one of the allowed values")
	ErrInvalidValue     = errors.New("invalid value provided")
)

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := fieldErr.Field()

		switch fieldErr.Tag() {
		case "required":
			if field == "UserID" {
				errors["user_id"] = ErrMissingUserID.Error()
			} else if field == "Date" {
				errors["date"] = ErrMissingDate.Error()
			} else {
				errors[field] = "This field is required"
			}
		case "uuid4":
			errors["user_id"] = ErrInvalidUUID.Error()
		case "datetime":
			errors["date"] = ErrInvalidDate.Error()
		case "oneof":
			errors[field] = ErrInvalidChoice.Error()
		default:
			errors[field] = ErrInvalidValue.Error()
		}
	}

	return errors
}

// 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusInternalServerError, map[string]string{"error": "The server encountered a problem"})
}

// 400 Bad Request
func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		formattedErrors := FormatValidationErrors(validationErrs)
		WriteJSONError(w, http.StatusBadRequest, formattedErrors)
		return
	}

	WriteJSONError(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

// 404 Not Found
func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}

// 422 Unprocessable Entity
func UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Unprocessable Entity: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
}
