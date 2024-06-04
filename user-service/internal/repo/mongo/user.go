package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// RegisterUser ...
func (dao *dbClient) RegisterUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, literals.Error) {
	coll := GetCollection(dao.db, literals.UsersCollection)

	errs := validateNewUserDetails(dao, ctx, u)
	if len(errs) > 0 {
		return nil, errs
	}

	res, err := coll.InsertOne(ctx, bson.D{
		{"firstName", u.FirstName},
		{"lastName", u.LastName},
		{"email", u.Email},
		{"phoneNumber", u.Username},
		{"password", u.Password},
		{"createdAt", time.Now()},
		{"updatedAt", nil},
	})

	if err != nil {
		errs["error"] = UserInsertionError.Error()
		return nil, errs
	}

	u.ID = res.InsertedID.(primitive.ObjectID)
	dao.log.Infof("RegisterUser:DBSuccess: a new user with ID: %s was inserted successfully. Data: %v", res.InsertedID, u)

	return u, nil
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

// FindUserByID finds a record matching email
func (dao *dbClient) FindUserByID(ctx context.Context, id string) (*UserSchemaType, error) {
	coll := GetCollection(dao.db, literals.UsersCollection)
	filter := bson.D{{"_id", id}}

	var u UserSchemaType
	err := coll.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
