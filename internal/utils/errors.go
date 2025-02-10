package utils

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInvalidUUID      = errors.New("invalid UUID format")
	ErrInvalidDate      = errors.New("invalid date format, expected YYYY-MM-DD")
	ErrMissingUserID    = errors.New("user_id is required")
	ErrInvalidParameter = errors.New("invalid parameter provided")
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// will implement proper logging
	log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// will implement proper logging
	log.Printf("Bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// will implement proper logging
	log.Printf("Not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusNotFound, "Not found")
}
