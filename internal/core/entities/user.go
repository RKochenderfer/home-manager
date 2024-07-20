package entities

import guards "home-manager/server/internal/core/shared"

type User struct {
	id          int32
	name        string
	totalPoints int32
}

func NewUser(id int32, name string, totalPoints int32) (User, error) {
	err := guards.GuardAgainstZeroNegative(id)
	if err != nil {
		return User{}, err
	}
	err = guards.GuardAgainstEmptyOrWhitespace(name)
	if err != nil {
		return User{}, err
	}

	return User{id, name, totalPoints}, nil
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
