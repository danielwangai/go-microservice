package literals

import "errors"

// Error is the structure of an error/errors
// type Error map[string]error
type Error map[string]string

var (
	InvalidRegisterUserRequestPayload = errors.New("invalid request payload for register user")
)
