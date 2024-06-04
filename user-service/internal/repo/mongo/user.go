package repo

import (
	"context"
	"errors"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	dao.log.Infof("Begin:FindUserByID ID: %s", id)
	coll := GetCollection(dao.db, literals.UsersCollection)
	// convert id string to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}

	var u UserSchemaType
	err = coll.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		dao.log.Infof("Begin:FindUserByID ERROR: %s", err)
		return nil, err
	}

	return &u, nil
}

// FollowUser allows user of id1 to follow user of id 2
func (dao *dbClient) FollowUser(ctx context.Context, follower, followed *UserSchemaType) (*UserFollowerSchemaType, error) {
	coll := GetCollection(dao.db, literals.UserFollowCollection)

	isFollowing, err := dao.IsFollowingUser(ctx, follower, followed)
	if err != nil {
		return nil, err
	}
	if isFollowing {
		return nil, errors.New("User is already followed")
	}

	createdAt := time.Now()
	// follow user
	res, err := coll.InsertOne(ctx, bson.D{
		{"follower", follower},
		{"followed", followed},
		{"createdAt", createdAt},
	})

	if err != nil {
		return nil, err
	}

	return &UserFollowerSchemaType{
		ID:        res.InsertedID.(primitive.ObjectID),
		Follower:  follower,
		Followed:  followed,
		CreatedAt: createdAt,
	}, nil
}

func (dao *dbClient) IsFollowingUser(ctx context.Context, follower, followed *UserSchemaType) (bool, error) {
	coll := GetCollection(dao.db, literals.UserFollowCollection)
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"follower._id", bson.D{{"$eq", follower.ID}}}},
				bson.D{{"followed._id", bson.D{{"$eq", followed.ID}}}},
			}},
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}
	var results []*UserFollowerSchemaType
	if err = cursor.All(context.TODO(), &results); err != nil {
		return false, err
	}

	dao.log.Infof("results: %+v", len(results))

	return len(results) > 0, nil
}
