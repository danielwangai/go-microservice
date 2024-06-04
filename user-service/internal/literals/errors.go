package literals

import "errors"

// Error is the structure of an error/errors
// type Error map[string]error
type Error map[string]string

var (
	InvalidRegisterUserRequestPayload = errors.New("invalid request payload for register user")
	InvalidLoginRequestPayload        = errors.New("invalid user login request payload for register user")
	LoginAttemptFailed                = errors.New("an error occurred when logging in. Try again later")
	InvalidLoginCredentials           = errors.New("invalid login credentials")
)
