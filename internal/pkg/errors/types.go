package errors

import "net/http"

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(reason, message string) *Error {
	return New(http.StatusBadRequest, reason, message)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
func IsBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(reason, message string) *Error {
	return New(http.StatusUnauthorized, reason, message)
}

// IsUnauthorized determines if err is an error which indicates a Unauthorized error.
func IsUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string) *Error {
	return New(http.StatusForbidden, reason, message)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
func IsForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(reason, message string) *Error {
	return New(http.StatusNotFound, reason, message)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
func IsNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(reason, message string) *Error {
	return New(http.StatusConflict, reason, message)
}

// IsConflict determines if err is an error which indicates a Conflict error.
func IsConflict(err error) bool {
	return Code(err) == http.StatusConflict
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string) *Error {
	return New(http.StatusInternalServerError, reason, message)
}

// IsInternalServer determines if err is an error which indicates an InternalServer error.
func IsInternalServer(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(reason, message string) *Error {
	return New(http.StatusServiceUnavailable, reason, message)
}

// IsServiceUnavailable determines if err is an error which indicates a ServiceUnavailable error.
func IsServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

// InvalidArgument new InvalidArgument error that is mapped to a 400 response.
func InvalidArgument(message string) *Error {
	return BadRequest("invalid_argument", message)
}

// IsInvalidArgument determines if err is an error which indicates a InvalidArgument error.
func IsInvalidArgument(err error) bool {
	return Reason(err) == "invalid_argument" && Code(err) == http.StatusBadRequest
}

// AlreadyExists new AlreadyExists error that is mapped to a 409 response.
func AlreadyExists(message string) *Error {
	return Conflict("already_exists", message)
}

// IsAlreadyExists determines if err is an error which indicates a AlreadyExists error.
func IsAlreadyExists(err error) bool {
	return Reason(err) == "already_exists" && Code(err) == http.StatusConflict
}

// DatabaseError new DatabaseError error that is mapped to a 500 response.
func DatabaseError(message string) *Error {
	return InternalServer("database_error", message)
}

// IsDatabaseError determines if err is an error which indicates a DatabaseError error.
func IsDatabaseError(err error) bool {
	return Reason(err) == "database_error" && Code(err) == http.StatusInternalServerError
}

// UnknownError new UnknownError error that is mapped to a 500 response.
func UnknownError(message string) *Error {
	return InternalServer("unknown_error", message)
}

// IsUnknownError determines if err is an error which indicates a UnknownError error.
func IsUnknownError(err error) bool {
	return Reason(err) == "unknown_error" && Code(err) == http.StatusInternalServerError
}
