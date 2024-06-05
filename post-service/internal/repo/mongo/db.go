package repo

import (
	"context"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbClient struct {
	db  *mongo.Database
	log *logrus.Logger
}

func New(db *mongo.Database, log *logrus.Logger) DAO {
	return &dbClient{db, log}
}

// NewMongoClient creates a new database instance
func NewMongoClient(ctx context.Context, log *logrus.Logger, dbUrl string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dbUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithError(err).Error("Failed to connect to database")
		return nil, err
	}

	log.Info("\nDatabase Connection successful\n")
	return client, nil
}

// Ping Mongo checks if there's a successful connection to mongodb
func PingMongo(log *logrus.Logger, client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.WithError(err).Error("Failed to ping mongo")
		return err
	}

	log.Info("Connected to MongoDB successfully!")
	return nil
}

// GetCollection gets specified collection from the database
func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
