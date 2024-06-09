package svc

import (
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"log"
)

func validateCreatePostInputs(u *PostServiceRequestType) literals.Error {
	log.Printf("Begin valation: svc layer")
	errs := literals.Error{}

	if u.Title == "" {
		errs["title"] = literals.PostTitleRequiredError.Error()
	}
	if u.Content == "" {
		errs["content"] = literals.PostContentRequiredError.Error()
	}
	if u.CreatorID == "" {
		errs["creatorId"] = literals.PostCreatorIDRequiredError.Error()
	}

	return errs
}

func validateAddCommentInputs(c *CommentServiceRequestType) literals.Error {
	errs := literals.Error{}

	if c.Comment == "" {
		errs["comment"] = literals.CommentRequiredError.Error()
	}
	if c.PostID == "" {
		errs["postId"] = literals.CommentPostIDRequiredError.Error()
	}
	if c.CommenterID == "" {
		errs["commenterId"] = literals.CommenterIDRequiredError.Error()
	}

	return errs
}
