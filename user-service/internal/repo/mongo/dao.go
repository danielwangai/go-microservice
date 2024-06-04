package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	RegisterUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, literals.Error)
}
