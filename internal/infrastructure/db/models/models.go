package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"index"`
	TotalPoints uint

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
	gorm.Model
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
