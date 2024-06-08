package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dao *dbClient) FindUserByID(ctx context.Context, id string) (*UserSchemaType, error) {
	coll := GetCollection(dao.db, literals.UsersCollection)
	filter := bson.M{"_id": id}

	var u UserSchemaType
	err := coll.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			dao.log.WithError(err).Errorf("could not find user by ID: %s", id)
			return nil, literals.UserDoesNotExist
		}
		dao.log.WithError(err).Errorf("an error ocurred when finding user by ID: %s", id)
		return nil, err
	}

	return &u, nil
}
