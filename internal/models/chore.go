package models

type Chore struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Points int32 `json:"points"`
}