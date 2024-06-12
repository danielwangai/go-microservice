package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// IsCommentUnique - it's usecase is while consuming comments from kafka to prevent saving a comment more than once
func (dao *dbClient) IsCommentUnique(ctx context.Context, commentId, postId, commenterId string) (bool, error) {
	coll := GetCollection(dao.db, literals.CommentsCollection)
	// convert id string to object id
	commentOID, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return false, err
	}
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return false, err
	}
	commenterOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return false, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", bson.D{{"$eq", commentOID}}}},
				bson.D{{"commentedBy._id", bson.D{{"$eq", commenterOID}}}},
				bson.D{{"post._id", bson.D{{"$eq", postOID}}}},
			}},
	}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return false, err
	}

	var res []*CommentSchemaType
	if err = cursor.All(context.TODO(), &res); err != nil {
		return false, err
	}

	// if len(res) = 0; comment is unique, otherwise not
	return len(res) == 0, nil
}
