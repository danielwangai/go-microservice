package repo

import "errors"

var (
	DuplicateEmailError    = errors.New("a user with similar email exists")
	DuplicateUsernameError = errors.New("a user with similar username exists")
	UserInsertionError     = errors.New("an error occurred while creating user")
)
