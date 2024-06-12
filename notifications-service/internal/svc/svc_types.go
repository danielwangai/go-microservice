package svc

import "time"

type CommentServiceType struct {
	ID        string
	Title     string
	Comment   string
	Post      *PostServiceType
	CreatedBy *UserServiceType
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

type PostServiceType struct {
	ID        string
	Title     string
	Content   string
	CreatedBy *UserServiceType
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

type UserServiceType struct {
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
