package repo

import (
	"context"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

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

	//isFollowing, err := dao.IsFollowingUser(ctx, follower, followed)
	//if err != nil {
	//	return nil, err
	//}
	//if isFollowing {
	//	return nil, errors.New("User is already followed")
	//}

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

func (dao *dbClient) AddUser(ctx context.Context, u *UserSchemaType) (*UserSchemaType, error) {
	coll := GetCollection(dao.db, literals.UsersCollection)

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
		return nil, literals.UserInsertionError
	}

	dao.log.Infof("a new user with ID: %s was inserted successfully. Data: %v", res.InsertedID, u)
	return u, nil
}
