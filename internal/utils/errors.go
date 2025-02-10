package utils

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInvalidUUID      = errors.New("invalid UUID format")
	ErrInvalidDate      = errors.New("invalid date: ensure the format is YYYY-MM-DD and values are realistic")
	ErrMissingUserID    = errors.New("user_id is required")
	ErrMissingDate      = errors.New("date is required")
	ErrInvalidParameter = errors.New("invalid parameter provided")
	ErrInvalidDateRange = errors.New("invalid date range")
)

// 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

// 400 Bad Request
func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

// 404 Not Found
func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusNotFound, "Not found")
}

// 422 Unprocessable Entity
func UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Unprocessable Entity: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusUnprocessableEntity, err.Error())
}
