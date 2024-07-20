package entities

import guards "home-manager/server/internal/core/shared"

type Chore struct {
	id          int32
	name        string
	description string
	points      int32
}

func (c *Chore) Id() int32 {
	return c.id
}

func (c *Chore) Name() string {
	return c.name
}

func (c *Chore) Description() string {
	return c.description
}

func (c *Chore) Points() int32 {
	return c.points
}

func NewChore(id int32, name string, description string, points int32) (*Chore, error) {
	if err := guards.GuardAgainstEmptyOrWhitespace(name); err != nil {
		return nil, err
	}

	if err := guards.GuardAgainstEmptyOrWhitespace(description); err != nil {
		return nil, err
	}

	if err := guards.GuardAgainstZeroNegative(points); err != nil {
		return nil, err
	}

	return &Chore{id, name, description, points}, nil
}
