package entities

import "home-manager/server/internal/core/shared"

type Room struct {
	id   int32
	name string
}

func (r *Room)GetId() int32 {
	return r.id
}

func (r *Room)GetName() string {
	return r.name
}


func NewRoom(id int32, name string) (Room, error) {
	if err := shared.GuardAgainstZeroNegative(id); err != nil {
		return Room{}, err
	}

	if err := shared.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return Room{}, err
	}

	return Room{id, name}, nil
}
