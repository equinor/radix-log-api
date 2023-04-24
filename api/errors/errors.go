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
}

func (e *StatusError) Status() Status {
	return e.APIStatus
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s: %s", e.APIStatus.Reason, e.APIStatus.Message)
}

func NewUnauthorizedError() *StatusError {
	return &StatusError{
		APIStatus: Status{
			Code:    http.StatusUnauthorized,
			Reason:  "Unauthorized",
			Message: "Unauthorized",
		},
	}
}

func NewForbiddenError() *StatusError {
	return &StatusError{
		APIStatus: Status{
			Code:    http.StatusForbidden,
			Reason:  "Forbidden",
			Message: "Access to resource is denied",
		},
	}
}

func NewInternalServerError() *StatusError {
	return &StatusError{
		APIStatus: Status{
			Code:    http.StatusInternalServerError,
			Reason:  "InternalServerError",
			Message: "Internal Server Error",
		},
	}
}
