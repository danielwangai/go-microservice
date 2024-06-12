package repo

import "context"

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, error)
	IsCommentUnique(ctx context.Context, commentId, postId, commenterId string) (bool, error)
	CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, error)
	FindPostByTitleAndCreator(ctx context.Context, title, creatorId string) (*PostSchemaType, error)
}
