package literals

import "errors"

// Error is the structure of an error/errors
type Error map[string]string

var (
	DuplicateCommentError                      = errors.New("this comment has already been saved")
	PostDBInsertionError                       = errors.New("an error occurred while inserting post to the database")
	PostMatchingTitleAndCreatorIdNotFound      = errors.New("post matching title and creator id not found")
	PostMatchingTitleAndCreatorIDAlreadyExists = errors.New("post matching title and creator id already exists")
	ObjectToByteArrayConversionError           = errors.New("could not convert response object to byte array")
	FailedToPublishMessageToKafka              = errors.New("an error occurred when publishing message to kafka")
	UserInsertionError                         = errors.New("an error occurred while creating user")
)
