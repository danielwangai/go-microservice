package svc

import "time"

type CommentServiceRequestType struct {
	ID        string
	Title     string
	Comment   string
	Post      *PostServiceRequestType
	CreatedBy *UserRequestServiceType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentServiceResponseType struct {
	ID        string
	Title     string
	Comment   string
	Post      *PostServiceResponseType
	CreatedBy *UserServiceResponseType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostServiceRequestType struct {
	ID        string
	Title     string
	Content   string
	CreatedBy *UserRequestServiceType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostServiceResponseType struct {
	ID        string
	Title     string
	Content   string
	CreatedBy *UserServiceResponseType
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequestServiceType struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
