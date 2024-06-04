package svc

import "errors"

var (
	// user operation errors
	UserFirstNameRequiredError = errors.New("user's first name is required")
	UserLastNameRequiredError  = errors.New("user's last name is required")
	UserEmailRequiredError     = errors.New("user's email is required")
	UserPasswordRequiredError  = errors.New("user's password is required")
	DuplicateEmailError        = errors.New("duplicate email error")
	DuplicatePhoneNumberError  = errors.New("duplicate phone number error")

	PasswordHashingError = errors.New("error hashing password")
)
