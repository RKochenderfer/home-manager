package models

import "time"

type Room struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:uniqueIndex`
}

type Chore struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"index"`
	Description string
	Points      uint   `gorm:"index"`
	Room        Room   `gorm:"foreignKey:idx_room"`
}

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"index"`
	TotalPoints uint
}

type Assignment struct {
	Id            uint      `gorm:"primaryKey"`
	AssignedTo    User      `gorm:"foreignKey:idx_user"`
	ChoreAssigned Chore     `gorm:"foreignKey:idx_chore"`
	Scalar        string
	DueDate       time.Time `gorm:"index"`
}

type Reward struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"index"`
	Description string
	Cost        uint   `gorm:"index"`
}

type Redemption struct {
	Id            uint       `gorm:"primaryKey"`
	User          User       `gorm:"foreignKey:idx_user"`
	Reward        Reward     `gorm:"foreignKey:idx_reward"`
	RedemptionDate time.Time `gorm:"index"`
}
