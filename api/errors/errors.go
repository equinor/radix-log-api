package errors

import (
	"fmt"
	"net/http"
)

type Status struct {
	Code    int32  `json:"code,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type APIStatus interface {
	Status() Status
}

type StatusError struct {
	APIStatus Status
	Cause     error
}

type ErrorOptions func(e *StatusError)

func WithCause(err error) ErrorOptions {
	return func(e *StatusError) {
		e.Cause = err
	}
}

func (e *StatusError) Status() Status {
	return e.APIStatus
}

func (e *StatusError) Error() string {
	errMsg := fmt.Sprintf("reason: %s, message: %s", e.APIStatus.Reason, e.APIStatus.Message)
	if e.Cause != nil {
		errMsg += fmt.Sprintf(", cause: %v", e.Cause)
	}
	return errMsg
}

func applyOptions(err *StatusError, options ...ErrorOptions) {
	for _, option := range options {
		option(err)
	}
}

func NewBadRequestError(options ...ErrorOptions) *StatusError {
	err := &StatusError{
		APIStatus: Status{
			Code:    http.StatusBadRequest,
			Reason:  "BadRequest",
			Message: "Bad request",
		},
	}
	applyOptions(err, options...)
	return err
}

func NewUnauthorizedError(options ...ErrorOptions) *StatusError {
	err := &StatusError{
		APIStatus: Status{
			Code:    http.StatusUnauthorized,
			Reason:  "Unauthorized",
			Message: "Unauthorized",
		},
	}
	applyOptions(err, options...)
	return err
}

func NewForbiddenError(options ...ErrorOptions) *StatusError {
	err := &StatusError{
		APIStatus: Status{
			Code:    http.StatusForbidden,
			Reason:  "Forbidden",
			Message: "Access to resource is denied",
		},
	}
	applyOptions(err, options...)
	return err
}

func NewNotFoundError(resourceType, resourceName string, options ...ErrorOptions) *StatusError {
	err := &StatusError{
		APIStatus: Status{
			Code:    http.StatusNotFound,
			Reason:  "NotFound",
			Message: fmt.Sprintf("%s '%s' not found", resourceType, resourceName),
		},
	}
	applyOptions(err, options...)
	return err
}

func NewInternalServerError(options ...ErrorOptions) *StatusError {
	err := &StatusError{
		APIStatus: Status{
			Code:    http.StatusInternalServerError,
			Reason:  "InternalServerError",
			Message: "Internal Server Error",
		},
	}
	applyOptions(err, options...)
	return err
}
