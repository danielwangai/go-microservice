package svc

import (
	"context"
	"errors"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
)

type SVC struct {
	dao repo.DAO
	log *logrus.Logger
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger) Svc {
	return &SVC{dao: dao, log: log}
}

func (s *SVC) AddComment(ctx context.Context, c *CommentServiceRequestType) (*CommentServiceResponseType, error) {
	// ensure a comment is not consumed more than once
	ok, err := s.dao.IsCommentUnique(ctx, c.ID, c.Post.ID, c.CreatedBy.ID)
	// a case where there are no duplicate comments in db but there's an error
	if !ok && err != nil {
		return nil, errors.New("an error occurred while checking if comment is unique")
	}
	if ok {
		return nil, literals.DuplicateCommentError
	}

	commentModel := convertCommentSvcToModelType(c)
	// create post
	comment, err := s.dao.AddComment(ctx, commentModel)
	if err != nil {
		return nil, err
	}

	commentSvc := convertCommentResponseModelTypeToSvcType(comment)
	// send to kafka for notification purposes

	return commentSvc, nil
}

func (s *SVC) CreatePost(ctx context.Context, p *PostServiceRequestType) (*PostServiceResponseType, error) {
	// check if post title by creator is a duplicate
	res, err := s.dao.FindPostByTitleAndCreator(ctx, p.Title, p.CreatedBy.ID)
	if err != nil && !errors.Is(err, literals.PostMatchingTitleAndCreatorIdNotFound) {
		s.log.WithError(err).Errorf("an error occurred while finding post by title")
		return nil, errors.New("an error occurred while finding post by title")
	}
	if res != nil {
		return nil, literals.PostMatchingTitleAndCreatorIDAlreadyExists
	}

	// convert post service type to model type
	pModel := convertPostRequestSvcTypeToModelType(p)

	// create post
	post, err := s.dao.CreatePost(ctx, pModel)
	if err != nil {
		return nil, err
	}

	svcPost := convertPostResponseModelTypeToSvcType(post)

	return svcPost, nil
}

func (s *SVC) AddUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, error) {
	// convert user service type to model layer type
	uModel := convertUserServiceRequestTypeToModelType(u)

	// save to db
	res, err := s.dao.AddUser(ctx, uModel)
	if err != nil {
		return nil, err
	}

	// convert from user model type to user response type for service layer
	uSvc := convertUserModelToUserServiceResponseType(res)

	return uSvc, nil
}

// StoreFollowInfo stores reccord of a user following another
// id1 is the user id of the follower
// id2 is the user id of the user being followed
func (s *SVC) StoreFollowInfo(ctx context.Context, id1, id2 string) (*UserFollowerServiceResponseType, error) {
	follower, err := s.dao.FindUserByID(ctx, id1)
	if err != nil {
		s.log.WithError(err).Errorf("user id of the follower not found: %s", id1)
		return nil, err
	}

	// check if user to be followed exists
	followed, err := s.dao.FindUserByID(ctx, id2)
	if err != nil {
		s.log.WithError(err).Errorf("user id of the user to be followed not found: %s", id2)
		return nil, err
	}

	// follow user
	followObj, err := s.dao.FollowUser(ctx, follower, followed)
	if err != nil {
		return nil, err
	}

	return convertUserFollowModelToServiceResponseType(followObj), nil
}
