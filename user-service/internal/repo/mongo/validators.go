package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

func validateNewUserDetails(dao *dbClient, ctx context.Context, u *UserSchemaType) literals.Error {
	errs := literals.Error{}

	// email must be unique
	if _, err := dao.FindUserByEmail(ctx, u.Email); err == nil {
		errs["email"] = DuplicateEmailError.Error()
	}
	// username must be unique
	if _, err := dao.FindUserByUsername(ctx, u.Username); err == nil {
		errs["username"] = DuplicateUsernameError.Error()
	}

	return errs
}
