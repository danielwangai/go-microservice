package svc

import "context"

type Svc interface {
	AddComment(ctx context.Context, c *CommentServiceRequestType) (*CommentServiceResponseType, error)
	CreatePost(ctx context.Context, p *PostServiceRequestType) (*PostServiceResponseType, error)
}
