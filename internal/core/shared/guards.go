package shared

import (
	"errors"
	"strings"

	"github.com/google/uuid"
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

func GuardAgainstZeroUuid(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("UUID cannot be the 0 UUID")
	}
	return nil
}
