package handlers

import (
	"github.com/danielwangai/twiga-foods/user-service/internal/svc"
)

func convertUserApiRequestTypeToSvcType(u *svc.UserAPIRequestType) *svc.UserServiceRequestType {
	return &svc.UserServiceRequestType{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
	}
}

func convertUserModelToUserAPIResponseType(u *svc.UserServiceResponseType) *svc.UserAPIResponseType {
	return &svc.UserAPIResponseType{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}
