package svc

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
)

type Svc interface {
	CreatePost(ctx context.Context, p *PostServiceRequestType) (*PostServiceResponseType, literals.Error)
	GetPosts(ctx context.Context) ([]*PostServiceResponseType, error)
	FindPostById(ctx context.Context, id string) (*PostServiceResponseType, error)
	AddComment(ctx context.Context, c *CommentServiceRequestType) (*CommentServiceResponseType, literals.Error)
	AddUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, literals.Error)
}
