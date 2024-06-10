package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
)

func validateNewUserDetails(dao *dbClient, ctx context.Context, u *UserSchemaType) literals.Error {
	errs := literals.Error{}

	// id must be unique
	if _, err := dao.FindUserByID(ctx, u.ID); err == nil {
		errs["email"] = literals.DuplicateUserIDError.Error()
	}
	// email must be unique
	if _, err := dao.FindUserByEmail(ctx, u.Email); err == nil {
		errs["email"] = literals.DuplicateEmailError.Error()
	}
	// username must be unique
	if _, err := dao.FindUserByUsername(ctx, u.Username); err == nil {
		errs["username"] = literals.DuplicateUsernameError.Error()
	}

	return errs
}
