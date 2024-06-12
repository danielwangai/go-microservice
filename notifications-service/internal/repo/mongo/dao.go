package repo

import "context"

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, error)
}
