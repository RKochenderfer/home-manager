package valueobjects

import (
	"fmt"
	"home-manager/server/internal/core/entities"
	"reflect"

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
	if err := guards.GuardAgainstEmptyOrWhitespace(scalar); err != nil {
		return Assignment{}, err
	}

	if err := GuardAgainstInvalidAssignmentStatus(assignmentStatus); err != nil {
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

var reflectedValues []string // lazy load in enum values

func GuardAgainstInvalidAssignmentStatus(as AssignmentStatus) error {
	v := getAssignmentStatusEnumValues()

	for _, s := range v {
		if s == string(as) { return nil }
	}

	return fmt.Errorf("assignment status %s, is not valid", as)
}

func getAssignmentStatusEnumValues() []string {
	if reflectedValues != nil { return reflectedValues }

	v := reflect.ValueOf(AssignmentStatusEnum)
	vals := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
        reflectedField := v.Field(i).Interface()
		as := reflectedField.(AssignmentStatus)
		vals[i] = string(as)
    }

	reflectedValues = vals
	return reflectedValues
}
