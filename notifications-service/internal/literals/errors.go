package literals

import "errors"

// Error is the structure of an error/errors
type Error map[string]string

var (
	DuplicateCommentError                      = errors.New("this comment has already been saved")
	PostDBInsertionError                       = errors.New("an error occurred while inserting post to the database")
	PostMatchingTitleAndCreatorIdNotFound      = errors.New("post matching title and creator id not found")
	PostMatchingTitleAndCreatorIDAlreadyExists = errors.New("post matching title and creator id already exists")
)
