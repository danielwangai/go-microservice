package repo

//
//func validateNewPostDetails(dao *dbClient, ctx context.Context, p *PostSchemaType) literals.Error {
//	errs := literals.Error{}
//
//	// post title must be unique
//	if _, err := dao.FindPostByTitleAndCreator(ctx, p.Title, p.CreatedBy); err == nil {
//		errs["email"] = literals.DuplicateEmailError.Error()
//	}
//
//	return errs
//}
