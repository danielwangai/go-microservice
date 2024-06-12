package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
)

// AddComment ...
func (dao *dbClient) AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, error) {
	coll := GetCollection(dao.db, literals.CommentsCollection)

	res, err := coll.InsertOne(ctx, bson.D{
		{"_id", c.ID},
		{"comment", c.Comment},
		{"commentedBy", c.CommentedBy},
		{"post", c.Post},
		{"createdAt", c.CreatedAt},
		{"updatedAt", nil},
	})

	if err != nil {
		return nil, err
	}

	dao.log.Infof("a new comment with ID: %s was inserted successfully. Data: %v", res.InsertedID, c)
	return c, nil
}
