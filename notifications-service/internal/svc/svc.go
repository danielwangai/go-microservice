package svc

import (
	"context"
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

func (s *SVC) AddComment(ctx context.Context, c *CommentServiceType) (*CommentServiceResponseType, error) {
	s.log.Infof("Comment: %v", c)

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
