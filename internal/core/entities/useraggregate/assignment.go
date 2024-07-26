package useraggregate

import (
	"errors"
	"fmt"
	"home-manager/server/internal/core/valueobjects"
	"reflect"

	guards "home-manager/server/internal/core/shared"

	"github.com/google/uuid"
)

type Assignment struct {
	id            uuid.UUID
	assignedTo    *User
	choreAssigned valueobjects.Chore
	scalar        string
	status        AssignmentStatus
}

func (a *Assignment) AssignedTo() *User {
	return a.assignedTo
}

func (a *Assignment) ChoreAssigned() valueobjects.Chore {
	return a.choreAssigned
}

func (a *Assignment) Scalar() string {
	return a.scalar
}

func (a *Assignment) Status() AssignmentStatus {
	return a.status
}

func NewAssignment(assignedTo *User, choreAssigned valueobjects.Chore, scalar string, assignmentStatus AssignmentStatus) (*Assignment, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(scalar); err != nil {
		return nil, err
	}

	if err := GuardAgainstInvalidAssignmentStatus(assignmentStatus); err != nil {
		return nil, err
	}

	return &Assignment{uuid.New(), assignedTo, choreAssigned, scalar, assignmentStatus}, nil
}

func AssignmentFrom(id uuid.UUID, choreAssigned valueobjects.Chore, scalar string, status AssignmentStatus) (*Assignment, error) {
	if err := guards.GuardAgainstZeroUuid(id); err != nil {
		return nil, err
	}
	if err := guards.GuardAgainstEmptyOrWhitespace(scalar); err != nil {
		return nil, err
	}

	if err := GuardAgainstInvalidAssignmentStatus(status); err != nil {
		return nil, err
	}

	return &Assignment{id: uuid.New(),choreAssigned: choreAssigned, scalar: scalar,status: status}, nil
}

func (a *Assignment) SetAssignedTo(user *User) error {
	if user == nil {
		return errors.New("assigned to user cannot be nil")
	}
	
	a.assignedTo = user
	return nil
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
		if s == string(as) {
			return nil
		}
	}

	return fmt.Errorf("assignment status %s, is not valid", as)
}

func getAssignmentStatusEnumValues() []string {
	if reflectedValues != nil {
		return reflectedValues
	}

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
