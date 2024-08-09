package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
	pb "home-manager/server/homemanager/gen"
)

func ToRoomsListResponse(rooms []*entities.Room) []*models.RoomResponse {
	var responseList []*models.RoomResponse

	for _, room := range rooms {
		mapped := models.RoomResponse{Id: room.GetId(), Name: room.GetName()}
		responseList = append(responseList, &mapped)
	}

	return responseList
}

func ToCreateRoomResponse(room *entities.Room) *pb.CreateRoomResponse {
	return &pb.CreateRoomResponse{
		Room: toRoom(room),
	}
}

func ToGetRoomResponse(room *entities.Room) *pb.GetRoomResponse {
	return &pb.GetRoomResponse{
		Room: toRoom(room),
	}
}

func ToGetAllRoomsResponse(rooms *[]entities.Room) *pb.GetAllRoomsResponse {
	var roomsList []*pb.Room
	for _, r := range *rooms {
		er := toRoom(&r)
		roomsList = append(roomsList, er)
	}

	return &pb.GetAllRoomsResponse{
		Rooms: roomsList,
	}
}

func toRoom(room *entities.Room) *pb.Room {
	return &pb.Room{
		Id: room.GetId(),
		Name: room.GetName(),
	}
}
