package repo

import "context"

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, error)
	IsCommentUnique(ctx context.Context, commentId, postId, commenterId string) (bool, error)
	CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, error)
	FindPostByTitleAndCreator(ctx context.Context, title, creatorId string) (*PostSchemaType, error)
	FindUserByID(ctx context.Context, id string) (*UserSchemaType, error)
	FollowUser(ctx context.Context, follower, followed *UserSchemaType) (*UserFollowerSchemaType, error)
	AddUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, error)
	GetFollowsByUserID(ctx context.Context, id string) ([]*UserFollowerSchemaType, error)
}
