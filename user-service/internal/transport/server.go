package transport

import (
	"context"
	"fmt"
	k "github.com/danielwangai/twiga-foods/user-service/internal/kafka"
	"github.com/danielwangai/twiga-foods/user-service/internal/svc"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/danielwangai/twiga-foods/user-service/internal/config"
	"github.com/danielwangai/twiga-foods/user-service/internal/logging"
	mgo "github.com/danielwangai/twiga-foods/user-service/internal/repo/mongo"
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

	// kafka
	producer, err := k.NewKafkaProducer(cfg.Kafka.BootstrapServers)
	if err != nil {
		log.WithError(err).Error("failed to create new kafka producer")
		return err
	}

	kafka := k.New(producer, cfg.Kafka.BootstrapServers, cfg.Kafka.Topic)

	service := svc.New(dao, log, kafka)
	server.Router.InitializeRoutes(ctx, service, log, dbClient)

	log.Infof("starting server on port %s", cfg.WebServer.Port)
	if err := http.ListenAndServe(":"+cfg.WebServer.Port, *server.Router); err != nil {
		log.WithError(err).Error("could not start the HTTP server")
		return err
	}

	return nil
}
