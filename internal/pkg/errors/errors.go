package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Error contains an error response from the server.
type Error struct {
	OrigErr  error             `json:"-"`
	Code     int               `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s", e.Code, e.Reason, e.Message)
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if target := new(Error); errors.As(err, &target) {
		return target.Code == e.Code
	}
	return false
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]string) *Error {
	err := *e
	err.Metadata = md
	return &err
}

// WithOrigErr with an original error info
func (e *Error) WithOrigErr(oe error) *Error {
	err := *e
	err.OrigErr = oe
	return &err
}

// New returns an error object for the code, message.
func New(code int, reason, message string) *Error {
	return &Error{
		Code:    code,
		Reason:  reason,
		Message: message,
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Code returns the code for a particular error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if target := new(Error); errors.As(err, &target) {
		return target.Code
	}
	return http.StatusInternalServerError
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if target := new(Error); errors.As(err, &target) {
		return target.Reason
	}
	return ""
}
