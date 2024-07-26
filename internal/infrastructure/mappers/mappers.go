package mappers

import (
	"fmt"
	"home-manager/server/internal/core/entities/useraggregate"
	"home-manager/server/internal/core/valueobjects"
	"home-manager/server/internal/infrastructure/db/models"
)

func MapModelsAssignmentToEntityAssignments(ma []models.Assignment) []useraggregate.Assignment {
	var uas []useraggregate.Assignment
	for _, a := range ma {
		chore, err := MapModelsChoreToValueObjectChore(a.ChoreAssigned)
		if err != nil {
			continue
		}
		newAssignment, err := useraggregate.AssignmentFrom(a.ID, chore, a.Scalar, useraggregate.AssignmentStatus(a.AssignmentStatus))
		if err != nil {
			println(fmt.Errorf("error creating assignment from database values: %s", err.Error()))
			continue
		}
		uas = append(uas, *newAssignment)
	}
	
	return uas
}

func MapModelsChoreToValueObjectChore(mc models.Chore) (valueobjects.Chore, error) {
	chore, err := valueobjects.ChoreFrom(mc.Name, mc.Instructions, int32(mc.Points), int32(mc.RoomID))
	if err != nil {
		println(fmt.Errorf("error creating assignment from database values: %s", err.Error()))
		return valueobjects.Chore{}, nil
	}

	return chore, nil
}
