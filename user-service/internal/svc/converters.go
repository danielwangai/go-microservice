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
