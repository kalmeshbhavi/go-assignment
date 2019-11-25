package errors

type Code interface {
	// ErrorCode returns an error code in string representation.
	ErrorCode() string
}

// StringCode represents an error Code in string.
type StringCode string

// ErrorCode implements the Code interface.
func (c StringCode) ErrorCode() string {
	return string(c)
}
