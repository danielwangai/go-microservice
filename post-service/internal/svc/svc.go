package svc

import (
	"context"
	"errors"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"github.com/danielwangai/twiga-foods/post-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
)

type SVC struct {
	dao repo.DAO
	log *logrus.Logger
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger) Svc {
	return &SVC{dao, log}
}

func (s *SVC) CreatePost(ctx context.Context, p *PostServiceRequestType) (*PostServiceResponseType, literals.Error) {
	// handle validation
	errs := validateCreatePostInputs(p)
	if len(errs) > 0 {
		return nil, errs
	}

	// check if creator exists in users collection
	u, err := s.dao.FindUserByID(ctx, p.CreatorID)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	// check if post title by creator is a duplicate
	res, err := s.dao.FindPostByTitleAndCreator(ctx, p.Title, p.CreatorID)
	if err != nil && !errors.Is(err, literals.PostMatchingTitleAndCreatorIdNotFound) {
		return nil, map[string]string{"error": err.Error()}
	}
	if res != nil {
		return nil, map[string]string{"error": "post matching title and creator id already exists"}
	}

	// convert post service type to model type
	pModel := convertPostRequestSvcTypeToModelType(p)
	pModel.CreatedBy = u

	// create post
	post, errs := s.dao.CreatePost(ctx, pModel)
	if errs != nil {
		return nil, errs
	}

	return convertPostResponseModelTypeToSvcType(post), nil
}
