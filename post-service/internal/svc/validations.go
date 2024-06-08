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

	log.Printf("Begin valation: svc layer: errors: %v", errs == nil)
	return errs
}
