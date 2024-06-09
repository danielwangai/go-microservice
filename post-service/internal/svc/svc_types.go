package svc

import (
	"time"
)

// PostServiceRequestType structure of a post request payload in the service layer
// after conversion from API layer
type PostServiceRequestType struct {
	Title     string
	Content   string
	CreatorID string
}

//type PostServicePayload

// PostServiceResponseType structure of a post after conversion from model layer
// after creation/update/retrieval
type PostServiceResponseType struct {
	ID        string
	Title     string
	Content   string
	CreatedBy *UserServiceResponseType
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PostAPIResponseType structure of a post in the API layer
type PostAPIResponseType struct {
	ID        string               `json:"id"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	CreatedBy *UserAPIResponseType `json:"createdBy"`
	CreatedAt string               `json:"createdAt"`
	UpdatedAt string               `json:"updatedAt"`
}

// PostAPIRequestType structure of a post in the API layer
type PostAPIRequestType struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatorID string `json:"creatorId"`
}

// UserServiceResponseType structure of a user after conversion from the model layer
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

// UserAPIResponseType structure of a user in the API layer
type UserAPIResponseType struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
