package svc

import repo "github.com/danielwangai/twiga-foods/post-service/internal/repo/mongo"

func convertPostRequestSvcTypeToModelType(p *PostServiceRequestType) *repo.PostSchemaType {
	return &repo.PostSchemaType{
		Title:   p.Title,
		Content: p.Content,
	}
}

func convertPostResponseModelTypeToSvcType(p *repo.PostSchemaType) *PostServiceResponseType {
	return &PostServiceResponseType{
		ID:        p.ID.Hex(),
		Title:     p.Title,
		Content:   p.Content,
		CreatedBy: convertUserModelResponseTypeToSvcType(p.CreatedBy),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func convertUserModelResponseTypeToSvcType(u *repo.UserSchemaType) *UserServiceResponseType {
	return &UserServiceResponseType{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  "", // hide password
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertCommentResponseModelTypeToSvcType(c *repo.CommentSchemaType) *CommentServiceResponseType {
	return &CommentServiceResponseType{
		ID:        c.ID.Hex(),
		Comment:   c.Comment,
		Post:      convertPostResponseModelTypeToSvcType(c.Post),
		CreatedBy: convertUserModelResponseTypeToSvcType(c.CommentedBy),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func convertUserServiceRequestTypeToModelType(u *UserServiceRequestType) *repo.UserSchemaType {
	return &repo.UserSchemaType{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
	}
}

func convertUserModelToUserServiceResponseType(u *repo.UserSchemaType) *UserServiceResponseType {
	return &UserServiceResponseType{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
