package user

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
