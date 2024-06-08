package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// CreatePost ...
func (dao *dbClient) CreatePost(ctx context.Context, p *PostSchemaType) (*PostSchemaType, literals.Error) {
	coll := GetCollection(dao.db, literals.PostsCollection)
	var errs literals.Error
	p.CreatedAt = time.Now()

	res, err := coll.InsertOne(ctx, bson.D{
		{"title", p.Title},
		{"content", p.Content},
		{"createdBy", p.CreatedBy},
		{"createdAt", p.CreatedAt},
		{"updatedAt", nil},
	})

	if err != nil {
		errs["error"] = literals.PostDBInsertionError.Error()
		return nil, errs
	}

	p.ID = res.InsertedID.(primitive.ObjectID)
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
