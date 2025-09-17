package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

// Errors
var (

	//General
	ErrQueryTimeOut        = errors.New("database query timed out")
	ErrRequestBodyTooLarge = errors.New("request body is too large")

	//ErrMissing
	ErrMissingUserID                  = errors.New("user_id is required")
	ErrMissingUsername                = errors.New("username is required")
	ErrMissingPolarID                 = errors.New("polar-id is required")
	ErrMissingOuraID                  = errors.New("oura-id is required")
	ErrMissingToken                   = errors.New("token is required")
	ErrMissingSector                  = errors.New("sector is required")
	ErrMissingDate                    = errors.New("date is required")
	ErrMissingGeneral                 = errors.New("this field is required")
	ErrMissingType                    = errors.New("type is required")
	ErrMissingSource                  = errors.New("source is required")
	ErrMissingHours                   = errors.New("hours is required")
	ErrMissingSportID                 = errors.New("sport_id is required")
	ErrMissingSporttiID               = errors.New("sportti_id is required")
	ErrMissingID                      = errors.New("id is required")
	ErrMissingStartTime               = errors.New("start_time is required")
	ErrMissingUpdatedAt               = errors.New("updated_at is required")
	ErrMissingCreatedAt               = errors.New("created_at is required")
	ErrMissingDuration                = errors.New("duration is required")
	ErrMissingSymptom                 = errors.New("symptom is required")
	ErrMissingSeverity                = errors.New("severity is required")
	ErrMissingName                    = errors.New("name is required")
	ErrMissingNameType                = errors.New("name_type is required")
	ErrMissingValue                   = errors.New("value is required")
	ErrMissingTimestamp               = errors.New("timestamp is required")
	ErrMissingData                    = errors.New("data is required")
	ErrMissingTypeID                  = errors.New("type_id is required")
	ErrMissingTypeResultType          = errors.New("type_result_type is required")
	ErrMissingQuestionnaireInstanceID = errors.New("questionnaire_instance_id is required")
	ErrMissingQuestionnaireKey        = errors.New("questionnaire_key is required")
	ErrMissingQuestionID              = errors.New("question_id is required")
	ErrMissingQuestionType            = errors.New("question_type is required")
	ErrMissingSessionID               = errors.New("session_id is required")
	ErrMissingRaceReport              = errors.New("race_report is required")
	ErrMissingMeasurementGroupID      = errors.New("measurement_group_id is required")
	ErrMissingNationalID              = errors.New("national_id is required")

	//ErrInvalid
	ErrInvalidUUID         = errors.New("invalid UUID")
	ErrInvalidDate         = errors.New("invalid date: ensure the format is YYYY-MM-DD and values are realistic")
	ErrInvalidTimeStamp    = errors.New("invalid timestamp: expected RFC3339. Examples: 2025-01-15T13:11:02Z, 2025-01-15T13:11:02+02:00, or 2025-01-15T13:11:02 (UTC assumed). Fractional seconds allowed")
	ErrInvalidParameter    = errors.New("invalid parameter provided")
	ErrInvalidDateRange    = errors.New("invalid date range")
	ErrInvalidChoice       = errors.New("invalid choice: must be one of the allowed values")
	ErrInvalidValue        = errors.New("invalid value provided")
	ErrInvalidSectorCode   = errors.New("invalid sector code. Allowed values: JP, NK, CC")
	ErrInvalidDevice       = errors.New("invalid device type. Allowed values: garmin, oura, polar, suunto")
	ErrInvalidSource       = errors.New("invalid source. Please use one of the allowed devices")
	ErrInvalidSportID      = errors.New("invalid sport_id format")
	ErrInvalidnumericValue = errors.New("value must be numeric")
	ErrInvalidIDNumeric    = errors.New("id must be numeric")
	ErrInvalidLimit        = errors.New("invalid limit: must be in the documented range")
	ErrInvalidOffset       = errors.New("invalid offset: must be a non-negative integer")
	ErrMaxLimitExceeded    = errors.New("maximum value exceeded: please use a smaller value")
	ErrMinLimitExceeded    = errors.New("minimum value not met: please use a larger value")

	// Database constraint errors
	ErrUserNotFound        = errors.New("user does not exist. Please create the user first")
	ErrExerciseNotFound    = errors.New("exercise does not exist")
	ErrForeignKeyViolation = errors.New("referenced record does not exist")
	ErrInvalidExerciseData = errors.New("exercise data contains invalid exercise id")
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
			case "PolarID":
				errors["polar-id"] = ErrMissingPolarID.Error()
			case "OuraID":
				errors["oura-id"] = ErrMissingOuraID.Error()
			case "Username":
				errors["username"] = ErrMissingUsername.Error()
			case "Token":
				errors["token"] = ErrMissingToken.Error()
			case "Date":
				errors["date"] = ErrMissingDate.Error()
			case "Type":
				errors["type"] = ErrMissingType.Error()
			case "Sector":
				errors["sector"] = ErrMissingSector.Error()
			case "Source":
				errors["source"] = ErrMissingSource.Error()
			case "Hours":
				errors["hours"] = ErrMissingHours.Error()
			case "SportID":
				errors["sport_id"] = ErrMissingSportID.Error()
			case "SporttiID":
				errors["sportti_id"] = ErrMissingSporttiID.Error()
			case "ID":
				errors["id"] = ErrMissingID.Error()
			case "StartTime":
				errors["start_time"] = ErrMissingStartTime.Error()
			case "UpdatedAt":
				errors["updated_at"] = ErrMissingUpdatedAt.Error()
			case "CreatedAt":
				errors["created_at"] = ErrMissingCreatedAt.Error()
			case "Duration":
				errors["duration"] = ErrMissingDuration.Error()
			case "Symptom":
				errors["symptom"] = ErrMissingSymptom.Error()
			case "Severity":
				errors["severity"] = ErrMissingSeverity.Error()
			case "Name":
				errors["name"] = ErrMissingName.Error()
			case "NameType":
				errors["name_type"] = ErrMissingNameType.Error()
			case "Value":
				errors["value"] = ErrMissingValue.Error()
			case "Timestamp":
				errors["timestamp"] = ErrMissingTimestamp.Error()
			case "Data":
				errors["data"] = ErrMissingData.Error()
			case "TypeID":
				errors["type_id"] = ErrMissingTypeID.Error()
			case "TypeResultType":
				errors["type_result_type"] = ErrMissingTypeResultType.Error()
			case "QuestionnaireInstanceID":
				errors["questionnaire_instance_id"] = ErrMissingQuestionnaireInstanceID.Error()
			case "QuestionnaireKey":
				errors["questionnaire_key"] = ErrMissingQuestionnaireKey.Error()
			case "QuestionID":
				errors["question_id"] = ErrMissingQuestionID.Error()
			case "QuestionType":
				errors["question_type"] = ErrMissingQuestionType.Error()
			case "SessionID":
				errors["session_id"] = ErrMissingSessionID.Error()
			case "RaceReport":
				errors["race_report"] = ErrMissingRaceReport.Error()
			case "MeasurementGroupID":
				errors["measurement_group_id"] = ErrMissingMeasurementGroupID.Error()
			case "NationalID":
				errors["national_id"] = ErrMissingNationalID.Error()
			default:
				errors[field] = ErrMissingGeneral.Error()
			}

		case "oneof":
			switch field {
			case "Sector":
				errors["sector"] = ErrInvalidSectorCode.Error()
			case "Device":
				errors["device"] = ErrInvalidDevice.Error()
			case "Source":
				errors["source"] = ErrInvalidSource.Error()
			default:
				errors[field] = ErrInvalidChoice.Error()
			}

		case "uuid4":
			switch field {
			case "UserID":
				errors["user_id"] = ErrInvalidUUID.Error()
			case "ID":
				errors["id"] = ErrInvalidUUID.Error()
			case "TypeID":
				errors["type_id"] = ErrInvalidUUID.Error()
			case "QuestionID":
				errors["question_id"] = ErrInvalidUUID.Error()
			case "QuestionnaireInstanceID":
				errors["questionnaire_instance_id"] = ErrInvalidUUID.Error()
			default:
				errors[field] = ErrInvalidUUID.Error()
			}

		case "datetime":
			errors["date"] = ErrInvalidDate.Error()

		case "numeric":
			switch field {
			case "SportID":
				errors["sport_id"] = ErrInvalidSportID.Error()
			case "SporttiID":
				errors["sportti_id"] = ErrInvalidSportID.Error()
			case "ID":
				errors["id"] = ErrInvalidIDNumeric.Error()
			case "NationalID":
				errors["national_id"] = ErrInvalidIDNumeric.Error()
			default:
				errors[field] = ErrInvalidnumericValue.Error()
			}

		case "max":
			errors[field] = ErrMaxLimitExceeded.Error()

		case "min":
			errors[field] = ErrMinLimitExceeded.Error()

		default:
			errors[field] = ErrInvalidValue.Error()
		}
	}

	return errors
}

// invalidTypeError
type InvalidFieldTypeError struct {
	Field        string
	ExpectedType string
	ActualType   string
}

func (e *InvalidFieldTypeError) Error() string {
	return fmt.Sprintf("value must be %s", e.ExpectedType)
}

func logError(r *http.Request, msg string, err error, status int) {
	requestID := middleware.GetReqID(r.Context())
	logger.Logger.Warnw(msg,
		"status", status,
		"error", err.Error(),
		"request_id", requestID,
	)
}

// 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, context.DeadlineExceeded) {
		logError(r, "Database timeout", err, http.StatusGatewayTimeout)
		WriteJSONError(w, http.StatusGatewayTimeout, map[string]string{"error": ErrQueryTimeOut.Error()})
		return
	}

	logError(r, "Internal server error", err, http.StatusInternalServerError)
	WriteJSONError(w, http.StatusInternalServerError, map[string]string{"error": "the server encountered a problem"})
}

// 400 Bad Request
func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil && strings.Contains(err.Error(), "request body too large") {
		logError(r, "Request body too large", err, http.StatusRequestEntityTooLarge)
		status := http.StatusRequestEntityTooLarge
		WriteJSONError(w, status, map[string]string{"error": ErrRequestBodyTooLarge.Error()})
		return
	}

	logError(r, "Bad request error", err, http.StatusBadRequest)

	// Handle validator validation errors
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		formattedErrors := FormatValidationErrors(validationErrs)
		WriteJSONError(w, http.StatusBadRequest, formattedErrors)
		return
	}

	// Handle JSON decoding errors
	if fieldErr, ok := err.(*InvalidFieldTypeError); ok {
		WriteJSONError(w, http.StatusBadRequest, map[string]string{
			toSnakeCase(fieldErr.Field): fieldErr.Error(),
		})
		return
	}

	WriteJSONError(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

// 404 Not Found
func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Not found error", err, http.StatusNotFound)
	WriteJSONError(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}

// 422 Unprocessable Entity
func UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Unprocessable Entity", err, http.StatusUnprocessableEntity)
	WriteJSONError(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
}

// 401 Unauthorized (JWT or client token)
func UnauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Unauthorized", err, http.StatusUnauthorized)
	WriteJSONError(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
}

// 401 Unauthorized (Basic Auth)
func UnauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Unauthorized (Basic Auth)", err, http.StatusUnauthorized)
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	WriteJSONError(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
}

// 403 Forbidden
func ForbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Forbidden", err, http.StatusForbidden)
	WriteJSONError(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
}

// 409 Conflict
func ConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, "Conflict", err, http.StatusConflict)
	WriteJSONError(w, http.StatusConflict, map[string]string{"error": err.Error()})
}

// 503 Service Unavailable for a specific database
func ServiceUnavailableDBResponse(w http.ResponseWriter, r *http.Request, dbName string) {
	err := fmt.Errorf("%s database is unavailable", dbName)
	logError(r, "Service unavailable", err, http.StatusServiceUnavailable)
	WriteJSONError(w, http.StatusServiceUnavailable, map[string]string{
		"error": err.Error(),
	})
}

// HandleDatabaseError analyzes database errors and returns appropriate HTTP responses - defualt 500 Internal Server Error
func HandleDatabaseError(w http.ResponseWriter, r *http.Request, err error) {
	// Check if it's a PostgreSQL error
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "22P02": // invalid_text_representation (includes JSON syntax errors)
			if strings.Contains(pqErr.Message, "json") || strings.Contains(pqErr.Message, "JSON") {
				// Try to identify which field has the JSON error
				if strings.Contains(pqErr.Message, "data") {
					BadRequestResponse(w, r, errors.New("invalid JSON format in 'data' field - check for missing braces, quotes, or commas"))
				} else if strings.Contains(pqErr.Message, "test_event_template_test_limits") {
					BadRequestResponse(w, r, errors.New("invalid JSON format in 'test_event_template_test_limits' field - check for missing braces, quotes, or commas"))
				} else if strings.Contains(pqErr.Message, "raw_data") {
					BadRequestResponse(w, r, errors.New("invalid JSON format in 'raw_data' field - check for missing braces, quotes, or commas"))
				} else if strings.Contains(pqErr.Message, "additional_info") {
					BadRequestResponse(w, r, errors.New("invalid JSON format in 'additional_info' field - check for missing braces, quotes, or commas"))
				} else if strings.Contains(pqErr.Message, "additional_data") {
					BadRequestResponse(w, r, errors.New("invalid JSON format in 'additional_data' field - check for missing braces, quotes, or commas"))
				} else {
					BadRequestResponse(w, r, errors.New("invalid JSON format in one of the JSON fields - check for missing braces, quotes, or commas"))
				}
				return
			}
			BadRequestResponse(w, r, fmt.Errorf("invalid data format: %s", pqErr.Message))
			return
		case "23503": // foreign_key_violation
			if strings.Contains(pqErr.Message, "user_id") || strings.Contains(pqErr.Detail, "user_id") {
				// Try to extract the specific user_id from the error detail
				if pqErr.Detail != "" {
					BadRequestResponse(w, r, fmt.Errorf("user does not exist. Details: %s", pqErr.Detail))
				} else {
					BadRequestResponse(w, r, ErrUserNotFound)
				}
				return
			} else if strings.Contains(pqErr.Message, "exercise_id") || strings.Contains(pqErr.Detail, "exercise_id") {
				// This usually means the main exercise insert was skipped due to conflict
				if pqErr.Detail != "" {
					BadRequestResponse(w, r, fmt.Errorf("exercise insert was skipped due to a conflict (duplicated raw_id, exercise_id, user_id for example), Details: %s", pqErr.Detail))
				} else {
					BadRequestResponse(w, r, ErrInvalidExerciseData)
				}
				return
			}
			// Generic foreign key violation with details if available
			if pqErr.Detail != "" {
				BadRequestResponse(w, r, fmt.Errorf("referenced record does not exist. Details: %s", pqErr.Detail))
			} else {
				BadRequestResponse(w, r, ErrForeignKeyViolation)
			}
			return
		case "23505": // unique_violation
			if pqErr.Detail != "" {
				ConflictResponse(w, r, fmt.Errorf("record already exists. Details: %s", pqErr.Detail))
			} else {
				ConflictResponse(w, r, errors.New("record already exists"))
			}
			return
		case "23514": // check_violation
			if pqErr.Detail != "" {
				BadRequestResponse(w, r, fmt.Errorf("data violates database constraints. Details: %s", pqErr.Detail))
			} else {
				BadRequestResponse(w, r, errors.New("data violates database constraints"))
			}
			return
		}
	}

	// Check for context timeout
	if errors.Is(err, context.DeadlineExceeded) {
		InternalServerError(w, r, ErrQueryTimeOut)
		return
	}

	// Default to internal server error
	InternalServerError(w, r, err)
}

// 429 Too Many Requests
func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	logError(r, "Rate limit", errors.New("rate limit exceeded"), http.StatusTooManyRequests)
	w.Header().Set("Retry-After", retryAfter)

	WriteJSONError(w, http.StatusTooManyRequests, map[string]string{
		"error":       "rate limit exceeded",
		"retry_after": retryAfter,
	})
}
