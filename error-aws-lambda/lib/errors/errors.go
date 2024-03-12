package errors

import (
	"fmt"
	"strings"
)

// TODO Add methods to access error fields

type ServiceError struct {
	cause  string
	code   int
	errMsg string
	detail []string
}

func New(msg string) *ServiceError {
	return &ServiceError{
		cause: msg,
	}
}

func (e *ServiceError) Error() string {
	// TODO improve for remove empty fields
	return fmt.Sprintf("cause: %s, code: %d, errMsg: %s, detail: %s", e.cause, e.code, e.errMsg, strings.Join(e.detail, ", "))
}

func (e *ServiceError) Code(status int) *ServiceError {
	e.code = status
	return e
}

func (e *ServiceError) Detail(detail ...string) *ServiceError {
	e.detail = append(e.detail, detail...)
	return e
}

func (e *ServiceError) ErrMsg(errMsg string) *ServiceError {
	e.errMsg = errMsg
	return e
}
