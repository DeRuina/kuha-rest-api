package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Errors
var (

	//General
	ErrQueryTimeOut = errors.New("database query timed out")

	//ErrMissing
	ErrMissingUserID  = errors.New("user_id is required")
	ErrMissingSector  = errors.New("sector is required")
	ErrMissingDate    = errors.New("date is required")
	ErrMissingGeneral = errors.New("this field is required")

	//ErrInvalid
	ErrInvalidUUID       = errors.New("invalid UUID")
	ErrInvalidDate       = errors.New("invalid date: ensure the format is YYYY-MM-DD and values are realistic")
	ErrInvalidParameter  = errors.New("invalid parameter provided")
	ErrInvalidDateRange  = errors.New("invalid date range")
	ErrInvalidChoice     = errors.New("invalid choice: must be one of the allowed values")
	ErrInvalidValue      = errors.New("invalid value provided")
	ErrInvalidSectorCode = errors.New("invalid sector code. Allowed values: JP, NK, CC")
)

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := fieldErr.Field()

		switch fieldErr.Tag() {
		case "required":
			switch field {
			case "UserID":
				errors["user_id"] = ErrMissingUserID.Error()
			case "Date":
				errors["date"] = ErrMissingDate.Error()
			case "Sector":
				errors["sector"] = ErrMissingSector.Error()
			default:
				errors[field] = ErrMissingGeneral.Error()
			}

		case "oneof":
			switch field {
			case "Sector":
				errors["sector"] = ErrInvalidSectorCode.Error()
			default:
				errors[field] = ErrInvalidChoice.Error()
			}

		case "uuid4":
			errors["user_id"] = ErrInvalidUUID.Error()

		case "datetime":
			errors["date"] = ErrInvalidDate.Error()

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
