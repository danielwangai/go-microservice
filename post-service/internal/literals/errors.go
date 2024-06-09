package literals

import "errors"

// Error is the structure of an error/errors
type Error map[string]string

var (
	PostDBInsertionError                  = errors.New("an error occurred while creating post")
	PostMatchingTitleAndCreatorIdNotFound = errors.New("post matching title and creator id not found")
	PostTitleRequiredError                = errors.New("post title cannot be empty")
	PostContentRequiredError              = errors.New("post content cannot be empty")
	PostCreatorIDRequiredError            = errors.New("post creator id must be provided")
	InvalidCreatePostRequestPayload       = errors.New("invalid request payload for register user")
	UserDoesNotExist                      = errors.New("user matching id does not exist")
)
