package svc

import "context"

type Svc interface {
	AddComment(ctx context.Context, c *CommentServiceRequestType) (*CommentServiceResponseType, error)
	CreatePost(ctx context.Context, p *PostServiceRequestType) (*PostServiceResponseType, error)
	StoreFollowInfo(ctx context.Context, id1, id2 string) (*UserFollowerServiceResponseType, error)
	AddUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, error)
}
