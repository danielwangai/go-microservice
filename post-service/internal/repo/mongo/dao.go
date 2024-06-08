package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
)

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
	CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, literals.Error)
	FindUserByID(ctx context.Context, id string) (*UserSchemaType, error)
	FindPostByTitleAndCreator(ctx context.Context, title, creatorId string) (*PostSchemaType, error)
}
