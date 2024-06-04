package svc

import (
	"time"
)

// UserServiceRequestType request payload of create/update operations in the service layer
// after conversion from API/http layer
type UserServiceRequestType struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

// UserServiceResponseType response payload of user operations in the service layer
// after conversion from db layer
type UserServiceResponseType struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserAPIRequestType json request payload for creating/updating a user
// from the API/http layer
type UserAPIRequestType struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// UserAPIResponseType json response payload for a user on create/update/fetch
type UserAPIResponseType struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
