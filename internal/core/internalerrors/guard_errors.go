package internalerrors

import "fmt"

type GuardError struct {
	guardType string
	text      string
	err       error
}

func (e *GuardError) Error() string {
	return fmt.Sprintf("%s %s %s", e.guardType, e.text, e.err.Error())
}

func New(guardType string, text string, err error) error {
	return &GuardError{guardType, text, err}
}
