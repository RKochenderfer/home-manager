package guards

import (
	"errors"
	"strings"
)

func GuardAgainstEmptyOrWhitespace(str string) error {
	if len(strings.TrimSpace(str)) == 0 {
		return errors.New("String cannot be empty or whitespace")
	}

	return nil
}

func GuardAgainstZeroNegative(num int32) error {
	if num < 0 {
		return errors.New("Number cannot be <= 0")
	}

	return nil
}
