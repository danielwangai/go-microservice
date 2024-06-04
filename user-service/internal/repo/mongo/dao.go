package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	RegisterUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, literals.Error)
	FindUserByID(ctx context.Context, id string) (*UserSchemaType, error)
	FindUserByUsername(ctx context.Context, username string) (*UserSchemaType, error)
	FindUserByEmail(ctx context.Context, email string) (*UserSchemaType, error)
}
