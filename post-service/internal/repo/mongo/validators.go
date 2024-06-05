package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
)

func validateNewPostDetails(dao *dbClient, ctx context.Context, p *PostSchemaType) literals.Error {
	errs := literals.Error{}

	// post title must be unique
	if _, err := dao.FindPostByTitle(ctx, p.Title); err == nil {
		errs["email"] = literals.DuplicateEmailError.Error()
	}

	return errs
}
