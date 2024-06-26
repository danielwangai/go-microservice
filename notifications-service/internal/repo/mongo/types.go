package repo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentSchemaType struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Comment     string             `bson:"comment"`
	Post        *PostSchemaType    `bson:"post"`
	CommentedBy *UserSchemaType    `bson:"commentedBy"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

type PostSchemaType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedBy *UserSchemaType    `bson:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserSchemaType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserFollowerSchemaType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Follower  *UserSchemaType    `bson:"follower"`
	Followed  *UserSchemaType    `bson:"followed"`
	CreatedAt time.Time          `bson:"createdAt"`
}
