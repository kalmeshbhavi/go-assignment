package errors

const (
	KindNotFound       Kind = "NotFound"
	KindUnexpected     Kind = "Unexpected"
	KindInvalidRequest Kind = "InvalidRequest"
	KindInternal       Kind = "Internal"
)

type Kind string

func KindOf(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.Kind != "" {
		return e.Kind
	}
	return KindOf(e.Err)
}
