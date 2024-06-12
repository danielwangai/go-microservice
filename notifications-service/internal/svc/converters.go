package svc

import (
	repo "github.com/danielwangai/twiga-foods/notifications-service/internal/repo/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func convertCommentSvcToModelType(c *CommentServiceRequestType) *repo.CommentSchemaType {
	// convert comment id string to object id
	objectID, _ := primitive.ObjectIDFromHex(c.ID)
	return &repo.CommentSchemaType{
		ID:          objectID,
		Comment:     c.Comment,
		Post:        convertPostSvcModelTypeToModelType(c.Post),
		CommentedBy: convertUserSvcTypeToModelType(c.CreatedBy),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
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

func convertUserSvcTypeToModelType(u *UserRequestServiceType) *repo.UserSchemaType {
	objectID, _ := primitive.ObjectIDFromHex(u.ID)
	return &repo.UserSchemaType{
		ID:        objectID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
	}
}

func convertUserModelResponseTypeToSvcType(u *repo.UserSchemaType) *UserServiceResponseType {
	return &UserServiceResponseType{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  "", // hide password
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertPostSvcModelTypeToModelType(p *PostServiceRequestType) *repo.PostSchemaType {
	objectID, _ := primitive.ObjectIDFromHex(p.ID)
	return &repo.PostSchemaType{
		ID:        objectID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedBy: convertUserSvcTypeToModelType(p.CreatedBy),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func convertPostRequestSvcTypeToModelType(p *PostServiceRequestType) *repo.PostSchemaType {
	objectID, _ := primitive.ObjectIDFromHex(p.ID)
	return &repo.PostSchemaType{
		ID:        objectID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedBy: convertUserSvcTypeToModelType(p.CreatedBy),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
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

func convertUserFollowModelToServiceResponseType(followObj *repo.UserFollowerSchemaType) *UserFollowerServiceResponseType {
	return &UserFollowerServiceResponseType{
		ID:        followObj.ID.Hex(),
		Follower:  convertUserModelToUserServiceResponseType(followObj.Follower),
		Followed:  convertUserModelToUserServiceResponseType(followObj.Followed),
		CreatedAt: followObj.CreatedAt.String(),
	}
}

func convertUserModelToUserServiceResponseType(u *repo.UserSchemaType) *UserServiceResponseType {
	return &UserServiceResponseType{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertUserServiceRequestTypeToModelType(u *UserServiceRequestType) *repo.UserSchemaType {
	objectID, _ := primitive.ObjectIDFromHex(u.ID)
	return &repo.UserSchemaType{
		ID:        objectID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
	}
}
