package valueobjects

import (
	"home-manager/server/internal/core/entities"

	guards "home-manager/server/internal/core/shared"
)

type Assignment struct {
	assignedTo    entities.User
	choreAssigned Chore
	scalar        string
	status        AssignmentStatus
}

// PartialEq implements ValueObject.
func (a Assignment) PartialEq(v ValueObject) bool {
	other, ok := v.(Assignment)
	if !ok {
		return false
	}

	return a.assignedTo == other.AssignedTo() &&
		a.choreAssigned == other.ChoreAssigned() &&
		a.scalar == other.Scalar() &&
		a.status == other.Status()
}

func (a *Assignment) AssignedTo() entities.User {
	return a.assignedTo
}

func (a *Assignment) ChoreAssigned() Chore {
	return a.choreAssigned
}

func (a *Assignment) Scalar() string {
	return a.scalar
}

func (a *Assignment) Status() AssignmentStatus {
	return a.status
}

func NewAssignment(assignedTo entities.User, choreAssigned Chore, scalar string, assignmentStatus AssignmentStatus) (ValueObject, error) {
	err := guards.GuardAgainstEmptyOrWhitespace(scalar)
	if err != nil {
		return Assignment{}, err
	}

	return Assignment{assignedTo, choreAssigned, scalar, assignmentStatus}, nil
}

type AssignmentStatus string

var AssignmentStatusEnum = struct {
	NotStarted     AssignmentStatus
	Started        AssignmentStatus
	ReadyForReview AssignmentStatus
	Completed      AssignmentStatus
	Canceled       AssignmentStatus
}{
	NotStarted:     "NotStarted",
	Started:        "Started",
	ReadyForReview: "ReadyForReview",
	Completed:      "Completed",
	Canceled:       "Canceled",
}
