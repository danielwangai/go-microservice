package svc

import "context"

type Svc interface {
	AddComment(ctx context.Context, c *CommentServiceType) (*CommentServiceResponseType, error)
}
