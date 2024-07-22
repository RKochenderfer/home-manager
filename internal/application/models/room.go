package models

type RoomResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type RoomRequest struct {
	Name string `json:"name"`
}