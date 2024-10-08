package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type User struct { 
	Base
	Name        string `gorm:"index"`
	TotalPoints uint
	Role        string `gorm:"size:20;index"`

	Assignments []Assignment `gorm:"foreignKey:ID"`
	Redemptions []Redemption `gorm:"foreignKey:ID"`
}

type Room struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`

	Chores []Chore `gorm:"constraint:OnDelete:CASCADE"`
}

type Chore struct {
	gorm.Model
	Name         string `gorm:"index"`
	Instructions string
	Points       uint `gorm:"index"`
	RoomID       uint
	Room         Room

	Assignments []Assignment `gorm:"foreignKey:ID"`
}

type Assignment struct {
	Base
	CreatedByUserID  uint `gorm:"index"`
	AssignedToUserID uint `gorm:"index"`
	ChoreAssignedID  uint `gorm:"index"`
	Scalar           string
	DueDate          time.Time `gorm:"index"`
	AssignmentStatus string `gorm:"default:NotStarted;size:20;index"`

	ChoreAssigned  Chore
	CreatedByUser  User
	AssignedToUser User
}

type Redemption struct {
	gorm.Model
	UserID         uint      `gorm:"index"`
	RedemptionDate time.Time `gorm:"index"`
	RewardID       uint

	User   User
	Reward Reward
}

type Reward struct {
	gorm.Model
	Name         string `gorm:"index"`
	Description  string
	Cost         uint `gorm:"index"`
	RedemptionID uint

	Redemption []Redemption
}
