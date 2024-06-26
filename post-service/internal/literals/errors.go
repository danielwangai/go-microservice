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
	PostMatchingIDNotFoundError           = errors.New("post matching ID not found")
	InvalidCreatePostRequestPayload       = errors.New("invalid request payload for register user")
	UserDoesNotExist                      = errors.New("user matching id does not exist")
	CommentDBInsertionError               = errors.New("an error occurred while commenting")
	CommentRequiredError                  = errors.New("comment cannot be empty")
	CommentPostIDRequiredError            = errors.New("post ID cannot be empty")
	CommenterIDRequiredError              = errors.New("comment id must be provided")
	InvalidAddCommentRequestPayload       = errors.New("invalid request payload for create comment")
	UserFirstNameRequiredError            = errors.New("user's first name is required")
	UserLastNameRequiredError             = errors.New("user's last name is required")
	UserEmailRequiredError                = errors.New("user's email is required")
	UserPasswordRequiredError             = errors.New("user's password is required")
	DuplicateUserIDError                  = errors.New("a user with similar ID already exists")
	UserInsertionError                    = errors.New("an error occurred while creating user")
	DuplicateEmailError                   = errors.New("a user with similar email exists")
	DuplicateUsernameError                = errors.New("a user with similar username exists")
)
