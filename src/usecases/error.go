package usecases

import "fmt"

type ErrWrongType struct {
	expected interface{}
	got      interface{}
}

func NewErrWrongType(expected interface{}, got interface{}) error {
	return &ErrWrongType{
		expected: expected,
		got:      got,
	}
}

func (err *ErrWrongType) Error() string {
	return fmt.Sprintf("wrong parameter type: expected %T; got: %T",
		err.expected, err.got)
}
