package errors

type Error struct {
	Op      Op
	Kind    Kind
	Err     error
	Context string
}

func (e *Error) Error() string {
	return e.Context
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		case string:
			e.Context = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

func (e *Error) WithContext(context string) {
	e.Context = context
}
