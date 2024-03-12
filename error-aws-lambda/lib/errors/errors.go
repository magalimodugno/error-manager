package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TODO Add methods to access error fields

type (
	ServiceError struct {
		cause  string
		code   int
		err    error
		detail []string
	}

	APIResponse struct {
		Cause  string   `json:"cause"`
		Code   int      `json:"httpStatus"`
		Detail []string `json:"detail"`
	}

	ErrorMessage struct { //nolint: errname
		Message string `json:"errorMessage"`
	}
)

func (e *ErrorMessage) Error() string {
	return e.Message
}

func New(msg string) *ServiceError {
	return &ServiceError{
		cause: msg,
	}
}

func (se *ServiceError) Error() string {
	// TODO improve for remove empty fields
	return fmt.Sprintf("cause: %s, code: %d, err %s, detail: %s", se.cause, se.code, se.err.Error(), strings.Join(se.detail, ", "))
}

func (se *ServiceError) Code(status int) *ServiceError {
	se.code = status
	return se
}

func (se *ServiceError) Detail(detail ...string) *ServiceError {
	se.detail = append(se.detail, detail...)
	return se
}

func (se *ServiceError) ErrMsg(err error) *ServiceError {
	se.err = err
	return se
}

func (se *ServiceError) MarshalJSON() ([]byte, error) {
	err := APIResponse{
		Cause:  se.cause,
		Code:   se.code,
		Detail: se.detail,
	}

	return json.Marshal(err)
}
