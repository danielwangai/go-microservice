package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// AddComment ...
func (dao *dbClient) AddComment(ctx context.Context, c *CommentSchemaType) (*CommentSchemaType, literals.Error) {
	coll := GetCollection(dao.db, literals.CommentsCollection)
	var errs literals.Error
	c.CreatedAt = time.Now()

	res, err := coll.InsertOne(ctx, bson.D{
		{"comment", c.Comment},
		{"commentedBy", c.CommentedBy},
		{"post", c.Post},
		{"createdAt", c.CreatedAt},
		{"updatedAt", nil},
	})

	if err != nil {
		errs["error"] = literals.CommentDBInsertionError.Error()
		return nil, errs
	}

	c.ID = res.InsertedID.(primitive.ObjectID)
	dao.log.Infof("a new post with ID: %s was inserted successfully. Data: %v", res.InsertedID, c)

	return c, nil
}
