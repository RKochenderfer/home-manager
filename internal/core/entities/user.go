package entities

import guards "home-manager/server/internal/core/shared"

type User struct {
	id          int32
	name        string
	totalPoints int32
	role        Role
}

func NewUser(id int32, name string, totalPoints int32, role Role) (User, error) {
	err := guards.GuardAgainstZeroNegative(id)
	if err != nil {
		return User{}, err
	}
	err = guards.GuardAgainstEmptyOrWhitespace(name)
	if err != nil {
		return User{}, err
	}

	return User{id, name, totalPoints, role}, nil
}

func (u *User) Id() int32 {
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

var RoleEnum = struct {
	Admin Role
	User  Role
}{
	Admin: "Admin",
	User: "User",
}
