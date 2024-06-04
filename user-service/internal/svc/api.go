package svc

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

type Svc interface {
	RegisterUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, literals.Error)
	FindUserByEmail(ctx context.Context, email string) (*UserServiceResponseType, error)
	FollowUser(ctx context.Context, id1, id2 string) (*UserFollowerServiceResponseType, error)
}
