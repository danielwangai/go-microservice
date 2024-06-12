package svc

import (
	"context"
	"encoding/json"
	"errors"
	k "github.com/danielwangai/twiga-foods/post-service/internal/kafka"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"github.com/danielwangai/twiga-foods/post-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
)

type SVC struct {
	dao   repo.DAO
	log   *logrus.Logger
	kafka *k.KafkaProducer
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger, kafkaProducer *k.KafkaProducer) Svc {
	return &SVC{dao, log, kafkaProducer}
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

	svcPost := convertPostResponseModelTypeToSvcType(post)

	// write to kafka that a new post has been created for notification purposes
	// convert post to bytes
	pBytes, err := json.Marshal(svcPost)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}
	err = s.kafka.PushMessageToQueue(literals.NewPostTopic, svcPost.ID, pBytes)
	if err != nil {
		s.log.WithError(err).Errorf("failed to push post notification to queue")
		return nil, map[string]string{"error": err.Error()}
	}

	return svcPost, nil
}

func (s *SVC) GetPosts(ctx context.Context) ([]*PostServiceResponseType, error) {
	posts, err := s.dao.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	var res []*PostServiceResponseType
	// convert posts list from model type to svc type
	for i := range posts {
		res = append(res, convertPostResponseModelTypeToSvcType(posts[i]))
	}

	return res, nil
}

func (s *SVC) FindPostById(ctx context.Context, id string) (*PostServiceResponseType, error) {
	user, err := s.dao.FindPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertPostResponseModelTypeToSvcType(user), nil
}

func (s *SVC) AddComment(ctx context.Context, c *CommentServiceRequestType) (*CommentServiceResponseType, literals.Error) {
	s.log.Infof("Comment: %v", c)
	// handle validation
	errs := validateAddCommentInputs(c)
	if len(errs) > 0 {
		return nil, errs
	}

	u, err := s.dao.FindUserByID(ctx, c.CommenterID)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	// check if post title by creator is a duplicate
	p, err := s.dao.FindPostByID(ctx, c.PostID)
	if err != nil && !errors.Is(err, literals.PostMatchingIDNotFoundError) {
		return nil, map[string]string{"error": err.Error()}
	}

	cModel := &repo.CommentSchemaType{
		Comment:     c.Comment,
		CommentedBy: u,
		Post:        p,
	}

	// create post
	comment, errs := s.dao.AddComment(ctx, cModel)
	if errs != nil {
		return nil, errs
	}
	commentSvc := convertCommentResponseModelTypeToSvcType(comment)
	// send to kafka for notification purposes

	// write to kafka that a new post has been created for notification purposes
	// convert post to bytes
	pBytes, err := json.Marshal(commentSvc)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}
	err = s.kafka.PushMessageToQueue(literals.NewCommentTopic, commentSvc.ID, pBytes)
	if err != nil {
		s.log.WithError(err).Errorf("failed to push new comment notification to queue")
		return nil, map[string]string{"error": err.Error()}
	}
	return commentSvc, nil
}

func (s *SVC) AddUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, literals.Error) {
	errs := validateRegisterUserInputs(u)
	if len(errs) > 0 {
		return nil, errs
	}

	// convert user service type to model layer type
	uModel := convertUserServiceRequestTypeToModelType(u)

	// save to db
	res, errs := s.dao.AddUser(ctx, uModel)
	if len(errs) > 0 {
		return nil, errs
	}

	// convert from user model type to user response type for service layer
	uSvc := convertUserModelToUserServiceResponseType(res)

	return uSvc, nil
}
