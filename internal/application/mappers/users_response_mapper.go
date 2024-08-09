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
	return &pb.CreateUserResponse{User: toProtoUser(user)}
}

func ToGetUserResponse(user *useraggregate.User) *pb.GetUserResponse {
	return &pb.GetUserResponse{User: toProtoUser(user)}
}

func ToGetAllUsersResponse(users *[]useraggregate.User) *pb.GetAllUsersResponse {
	var res []*pb.User
	for _, u := range *users {
		protoUser := toProtoUser(&u)
		res = append(res, protoUser)
	}

	return &pb.GetAllUsersResponse{Users: res}
}

func toProtoUser(user *useraggregate.User) *pb.User {
	return &pb.User{
		Id:          user.Id().String(),
		Name:        user.Name(),
		TotalPoints: user.TotalPoints(),
		Role:        string(user.Role()),
	}
}
