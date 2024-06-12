package literals

import "errors"

// Error is the structure of an error/errors
type Error map[string]string

var (
	DuplicateCommentError = errors.New("this comment has already been saved")
)
