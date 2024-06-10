package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser ...
func (dao *dbClient) AddUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, literals.Error) {
	coll := GetCollection(dao.db, literals.UsersCollection)
	errs := validateNewUserDetails(dao, ctx, u)
	if len(errs) > 0 {
		return nil, errs
	}

	res, err := coll.InsertOne(ctx, bson.D{
		{"_id", u.ID},
		{"firstName", u.FirstName},
		{"lastName", u.LastName},
		{"email", u.Email},
		{"username", u.Username},
		{"password", ""}, // no usecase of password in the post-service microservice
		{"createdAt", u.CreatedAt},
		{"updatedAt", nil},
	})

	if err != nil {
		errs["error"] = literals.UserInsertionError.Error()
		return nil, errs
	}

	dao.log.Infof("a new user with ID: %s was inserted successfully. Data: %v", res.InsertedID, u)
	return u, nil
}

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

// FindUserByEmail finds a record matching email
func (dao *dbClient) FindUserByEmail(ctx context.Context, email string) (*UserSchemaType, error) {
	coll := GetCollection(dao.db, literals.UsersCollection)
	filter := bson.D{{"email", email}}

	var u UserSchemaType
	err := coll.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// FindUserByUsername finds a record matching username
func (dao *dbClient) FindUserByUsername(ctx context.Context, username string) (*UserSchemaType, error) {
	coll := GetCollection(dao.db, literals.UsersCollection)
	filter := bson.D{{"username", username}}

	var u UserSchemaType
	err := coll.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
