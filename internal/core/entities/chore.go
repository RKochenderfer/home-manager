package entities

import guards "home-manager/server/internal/core/shared"

type Chore struct {
	id           int32
	name         string
	instructions string
	points       int32
	roomId       int32
}

func (c *Chore) Id() int32 {
	return c.id
}

func (c *Chore) Name() string {
	return c.name
}

func (c *Chore) Instructions() string {
	return c.instructions
}

func (c *Chore) Points() int32 {
	return c.points
}

func (c *Chore) RoomId() int32 {
	return c.roomId
}

func NewChore(id int32, name string, description string, points int32, roomId int32) (Chore, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return Chore{}, err
	}

	if err := guards.GuardAgainstEmptyOrWhitespace(description); err != nil {
		return Chore{}, err
	}

	if err := guards.GuardAgainstZeroNegative(points); err != nil {
		return Chore{}, err
	}

	return Chore{id, name, description, points, roomId}, nil
}

func From(name string, description string, points int32, roomId int32) (Chore, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return Chore{}, err
	}

	if err := guards.GuardAgainstEmptyOrWhitespace(description); err != nil {
		return Chore{}, err
	}

	if err := guards.GuardAgainstZeroNegative(points); err != nil {
		return Chore{}, err
	}

	return Chore{0, name, description, points, roomId}, nil
}
