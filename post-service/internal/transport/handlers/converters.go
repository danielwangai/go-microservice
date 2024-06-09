package handlers

import (
	"github.com/danielwangai/twiga-foods/post-service/internal/svc"
)

func convertCreatePostApiRequestTypeToSvcType(p *svc.PostAPIRequestType) *svc.PostServiceRequestType {
	return &svc.PostServiceRequestType{
		Title:     p.Title,
		Content:   p.Content,
		CreatorID: p.CreatorID,
	}
}

func convertPostSvcResponseTypeToAPIType(p *svc.PostServiceResponseType) *svc.PostAPIResponseType {
	return &svc.PostAPIResponseType{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedBy: convertUserSvcResponseTypeToAPIType(p.CreatedBy),
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func convertUserSvcResponseTypeToAPIType(p *svc.UserServiceResponseType) *svc.UserAPIResponseType {
	return &svc.UserAPIResponseType{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Username:  p.Username,
		Password:  "",
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func convertAddCommentAPIRequestTypeToSvcType(c *svc.CommentAPIRequestType) *svc.CommentServiceRequestType {
	return &svc.CommentServiceRequestType{
		Comment: c.Comment,
		//PostID:      c.PostID,
		CommenterID: c.CommenterID,
	}
}

func convertCommentSvcResponseTypeToAPIType(c *svc.CommentServiceResponseType) *svc.CommentAPIResponseType {
	return &svc.CommentAPIResponseType{
		ID:        c.ID,
		Title:     c.Title,
		Comment:   c.Comment,
		Post:      convertPostSvcResponseTypeToAPIType(c.Post),
		CreatedBy: convertUserSvcResponseTypeToAPIType(c.CreatedBy),
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
