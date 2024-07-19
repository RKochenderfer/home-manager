package assignment

import (
	"home-manager/server/internal/core/chore"
	guards "home-manager/server/internal/core/shared"
	"home-manager/server/internal/core/user"
)

type Assignment struct {
	id int32
	assignedTo user.User
	choreAssigned chore.Chore
	scalar string
}

func NewAssignment(id int32, assignedTo user.User, choreAssigned chore.Chore, scalar string) (Assignment, error) {
	err := guards.GuardAgainstZeroNegative(id)
	if err != nil {
		return Assignment{}, err
	}

	err = guards.GuardAgainstEmptyOrWhitespace(scalar)
	if err != nil {
		return Assignment{}, err
	}

	return Assignment{id, assignedTo, choreAssigned, scalar}, nil
}
