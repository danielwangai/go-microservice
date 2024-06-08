package repo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSchemaType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt"`
}

type PostSchemaType struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedBy *UserSchemaType    `bson:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt"`
}

type PostModelRequestType struct {
	Title   string
	Content string
	Creator *UserSchemaType
}
