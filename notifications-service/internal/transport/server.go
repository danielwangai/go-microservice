package transport

import (
	"context"
	"fmt"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/literals"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/svc"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/danielwangai/twiga-foods/notifications-service/internal/config"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/logging"
	mgo "github.com/danielwangai/twiga-foods/notifications-service/internal/repo/mongo"
)

// Server ...
type Server struct {
	Router *Router
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		Router: NewRouter(),
	}
}

// RunServer initializes services
func RunServer() error {
	log := logging.SetJSONFormatter(logrus.New())
	// gathers additional information about the env and
	ctx := context.Background()

	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Println("Error loading configs: ", err)
		return err
	}

	server := NewServer()

	// connect to mongodb
	dbClient, err := mgo.NewMongoClient(ctx, log, cfg.DB.DbURL)
	if err != nil {
		return err
	}

	// database
	db := dbClient.Database(cfg.DB.DbName)

	dao := mgo.New(db, log)

	// kafka producer
	kafkaTopicsMap := map[string]string{
		literals.NewUserTopic:    cfg.Kafka.Topics.NewUsersTopic,
		literals.NewPostTopic:    cfg.Kafka.Topics.NewPostNotificationTopic,
		literals.NewCommentTopic: cfg.Kafka.Topics.NewCommentNotificationTopic,
	}

	service := svc.New(dao, log)

	// kafka
	conn, err := connectConsumer([]string{cfg.Kafka.Broker})
	if err != nil {
		log.WithError(err).Error("Error connecting to consumer")
		return err
	}

	consumerConfig := NewKafkaConsumerConfig(conn, service, log)

	go consumerConfig.ConsumeNewComments([]string{cfg.Kafka.Broker}, kafkaTopicsMap[literals.NewCommentTopic])

	// initialize routes
	server.Router.InitializeRoutes(ctx, service, log, dbClient)

	log.Infof("starting server on port %s", cfg.WebServer.Port)
	if err := http.ListenAndServe(":"+cfg.WebServer.Port, *server.Router); err != nil {
		log.WithError(err).Error("could not start the HTTP server")
		return err
	}

	return nil
}
