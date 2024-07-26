package entities

import (
	"fmt"
	guards "home-manager/server/internal/core/shared"
	"reflect"

	"github.com/google/uuid"
)

type User struct {
	id          uuid.UUID
	name        string
	totalPoints int32
	role        Role
}

func NewUser(name string, totalPoints int32, role Role) (*User, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return nil, err
	}
	if err := GuardAgainstInvalidRole(role); err != nil {
		return nil, err
	}

	return &User{uuid.New(), name, totalPoints, role}, nil
}

func From(id uuid.UUID, name string, totalPoints int32, role Role) (*User, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return nil, err
	}
	if err := GuardAgainstInvalidRole(role); err != nil {
		return nil, err
	}

	return &User{id, name, totalPoints, role}, nil
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) TotalPoints() int32 {
	return u.totalPoints
}

func (u *User) Role() Role {
	return u.role
}

type Role string
var reflectValues []string

var RoleEnum = struct {
	Admin Role
	User  Role
}{
	Admin: "Admin",
	User: "User",
}

func GuardAgainstInvalidRole(r Role) error {
	v := getRoleEnumValues()

	for _, s := range v {
		if s == string(r) { return nil }
	}

	return fmt.Errorf("role %s, is not a valid role", r)
}

func getRoleEnumValues() []string {
	if reflectValues != nil {
		return reflectValues
	}

	v := reflect.ValueOf(RoleEnum)
	vals := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
        reflectedField := v.Field(i).Interface()
		role := reflectedField.(Role)
		vals[i] = string(role)
    }

	reflectValues = vals
	return reflectValues
}

