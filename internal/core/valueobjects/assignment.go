package valueobjects

import (
	"home-manager/server/internal/core/entities"

	guards "home-manager/server/internal/core/shared"
)

type Assignment struct {
	assignedTo    entities.User
	choreAssigned Chore
	scalar        string
}

func (a *Assignment) GetAssignedTo() entities.User {
	return a.assignedTo
}

func (a *Assignment) GetChoreAssigned() Chore {
	return a.choreAssigned
}

func (a *Assignment) GetScalar() string {
	return a.scalar
}

func NewAssignment(assignedTo entities.User, choreAssigned Chore, scalar string) (Assignment, error) {
	err := guards.GuardAgainstEmptyOrWhitespace(scalar)
	if err != nil {
		return Assignment{}, err
	}

	return Assignment{assignedTo, choreAssigned, scalar}, nil
}
