package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePost ...
func (dao *dbClient) CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, error) {
	coll := GetCollection(dao.db, literals.PostsCollection)
	res, err := coll.InsertOne(ctx, bson.D{
		{"_id", p.ID},
		{"title", p.Title},
		{"content", p.Content},
		{"createdBy", p.CreatedBy},
		{"createdAt", p.CreatedAt},
		{"updatedAt", nil},
	})

	if err != nil {
		return nil, literals.PostDBInsertionError
	}

	dao.log.Infof("a new post with ID: %s was inserted successfully. Data: %v", res.InsertedID, p)

	return p, nil
}

func (dao *dbClient) FindPostByTitleAndCreator(ctx context.Context, title, creatorId string) (*PostSchemaType, error) {
	coll := GetCollection(dao.db, literals.PostsCollection)
	// convert id string to object id
	objectID, err := primitive.ObjectIDFromHex(creatorId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"title", bson.D{{"$eq", title}}}},
				bson.D{{"createdBy._id", bson.D{{"$eq", objectID}}}},
			}},
	}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var res []*PostSchemaType
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, literals.PostMatchingTitleAndCreatorIdNotFound
	}

	return res[0], nil
}
