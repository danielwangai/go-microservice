package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
)

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	// user methods
	FindUserByID(ctx context.Context, id string) (*UserSchemaType, error)
	AddUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, literals.Error)
	// post methods
	CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, literals.Error)
	FindPostByTitleAndCreator(ctx context.Context, title, creatorId string) (*PostSchemaType, error)
	GetPosts(ctx context.Context) ([]*PostSchemaType, error)
	FindPostByID(ctx context.Context, id string) (*PostSchemaType, error)
	AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, literals.Error)
}
