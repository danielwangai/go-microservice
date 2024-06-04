package svc

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

type Svc interface {
	RegisterUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, literals.Error)
}
