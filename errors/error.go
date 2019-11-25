package errors

import (
	"fmt"
	"strings"
)

type ServiceError interface {
	GetCode() Code
	GetError() error

	ErrCodeIs(codes ...Code) bool

	Kind() string
	Message() string
	StackTrace() []string

	Error() string
}

type serviceError struct {
	code  Code
	error error
}

func (se *serviceError) GetCode() Code {
	return se.code
}

func (se *serviceError) GetError() error {
	return se.error
}

func (se *serviceError) Kind() string {
	return se.GetCode().ErrorCode()
}

func (se *serviceError) Message() string {
	errString := se.errString()

	return errString
}

func (se *serviceError) StackTrace() []string {
	//cs, _ := failure.CallStackOf(se.error)

	//return framesToString(cs.Frames())
	return []string{}
}

func (se *serviceError) Error() string {
	errString := se.errString()

	return fmt.Sprintf("code(%s): %s", se.code, errString)
}

func (se *serviceError) errString() string {
	errString := se.error.Error()
	errString = strings.TrimSpace(errString)
	errString = strings.TrimSuffix(errString, ":")

	return errString
}

func (se *serviceError) IsTransientError() bool {
	return se.ErrCodeIs(getTransientErrors()...)
}

func (se *serviceError) ErrCodeIs(codes ...Code) bool {
	return errCodeIs(se.GetCode(), codes...)
}

func NewFromError(err error, code Code, wrappers ...error) ServiceError {
	if isInvalid(code) {
		code = Unknown
	}

	return withCode(code, err)
}

func New(code Code, err error) ServiceError {
	if isInvalid(code) {
		code = Unknown
	}

	return withCode(code, err)
}
