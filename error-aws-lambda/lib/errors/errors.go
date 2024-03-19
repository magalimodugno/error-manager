package errors

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Bancar/uala-bis-go-dependencies/v2/constants/errors"
)

type (
	ServiceError struct {
		cause  string
		status *int
		err    error
		detail []string
	}

	Response struct {
		Cause  string   `json:"cause"`
		Status *int     `json:"httpStatus,omitempty"`
		Detail []string `json:"detail,omitempty"`
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

// TODO investigate better way to implement this
func (se *ServiceError) Error() string {
	msg := fmt.Sprintf("cause: %s", se.cause)

	if se.status != nil {
		msg = msg + fmt.Sprintf(", status: %d", se.status)
	}

	if se.err != nil {
		msg = msg + fmt.Sprintf(", err: %s", se.err.Error())
	}
	if len(se.detail) != 0 {
		msg = msg + fmt.Sprintf(", detail: %s", strings.Join(se.detail, ", "))
	}

	return msg
}

func (se *ServiceError) Status(status int) *ServiceError {
	se.status = &status
	return se
}

func (se *ServiceError) Detail(detail ...string) *ServiceError {
	se.detail = append(se.detail, detail...)
	return se
}

func (se *ServiceError) Err(err error) *ServiceError {
	se.err = err
	return se
}

func (se *ServiceError) MarshalJSON() ([]byte, error) {
	err := Response{
		Cause:  se.cause,
		Status: se.status,
		Detail: se.detail,
	}

	return json.Marshal(err)
}

func AsErrorMessage(err error) error {
	data, _ := json.Marshal(err)

	return &errors.ErrorMessage{
		Message: string(data),
	}
}

//func (se *ServiceError) StatusUnauthorized() *ServiceError {
//	se.status = http.StatusUnauthorized
//	return se
//}
//
//func (se *ServiceError) StatusForbidden() *ServiceError {
//	se.status = http.StatusForbidden
//	return se
//}
//
//func (se *ServiceError) StatusNotFound() *ServiceError {
//	se.status = http.StatusNotFound
//	return se
//}
//
//func (se *ServiceError) StatusInternalServerError() *ServiceError {
//	se.status = http.StatusInternalServerError
//	return se
//}
