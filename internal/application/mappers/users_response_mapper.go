package mappers

import (
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities/useraggregate"
)

func ToUserListResponse(users []useraggregate.User) []*models.UserResponse {
	var responseList []*models.UserResponse

	for _, user := range users {
		mapped := models.FromUser(&user)
		responseList = append(responseList, &mapped)
	}

	return responseList
}

func ToCreateUserResponse(user *useraggregate.User) *pb.CreateUserResponse {
	u := pb.User{
		Id:          user.Id().String(),
		Name:        user.Name(),
		TotalPoints: user.TotalPoints(),
		Role:        string(user.Role()),
	}

	return &pb.CreateUserResponse{User: &u}
}

func ToGetUserResponse(user *useraggregate.User) *pb.GetUserResponse {
	u := pb.User{
		Id:          user.Id().String(),
		Name:        user.Name(),
		TotalPoints: user.TotalPoints(),
		Role:        string(user.Role()),
	}

	return &pb.GetUserResponse{User: &u}
}
