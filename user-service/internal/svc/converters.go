package svc

import (
	repo "github.com/danielwangai/twiga-foods/user-service/internal/repo/mongo"
)

func convertUserServiceRequestTypeToModelType(u *UserServiceRequestType) *repo.UserSchemaType {
	return &repo.UserSchemaType{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
	}
}

func convertUserModelToUserServiceResponseType(u *repo.UserSchemaType) *UserServiceResponseType {
	return &UserServiceResponseType{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertUserFollowModelToServiceResponseType(followObj *repo.UserFollowerSchemaType) *UserFollowerServiceResponseType {
	return &UserFollowerServiceResponseType{
		ID:        followObj.ID.Hex(),
		Follower:  convertUserModelToUserServiceResponseType(followObj.Follower),
		Followed:  convertUserModelToUserServiceResponseType(followObj.Followed),
		CreatedAt: followObj.CreatedAt.String(),
	}
}
