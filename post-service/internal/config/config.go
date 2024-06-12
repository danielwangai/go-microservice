package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "POST_SERVICE"
)

// WebServerConfig ...
type WebServerConfig struct {
	Port string `envconfig:"POST_SERVICE_SERVER_PORT" split_words:"true"`
}

type MongoConfig struct {
	DbURL  string `envconfig:"POST_SERVICE_DATABASE_URL"`
	DbName string `envconfig:"POST_SERVICE_DATABASE_NAME"`
}

type KafkaConfig struct {
	NewUsersTopic            string `envconfig:"POST_SERVICE_KAFKA_NEW_USERS_TOPIC"`
	NewPostNotificationTopic string `envconfig:"POST_SERVICE_NEW_POST_NOTIFICATION_TOPIC"`
	GroupID                  string `envconfig:"POST_SERVICE_KAFKA_GROUP_ID"`
	Network                  string `envconfig:"POST_SERVICE_KAFKA_NETWORK"`
	Broker                   string `envconfig:"POST_SERVICE_KAFKA_BROKER"`
}

type AppConfig struct {
	WebServer *WebServerConfig
	DB        *MongoConfig
	Kafka     *KafkaConfig
}

// FromEnv loads the app config from environment variables
func FromEnv() (*AppConfig, error) {
	fromFileToEnv()
	cfg := &AppConfig{}
	if err := envconfig.Process(envPrefix, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() { // determine config file loc, irrespective of the entry (main or test); it should resolve properly
	cfgFilename := os.Getenv("ENV_FILE")
	if cfgFilename != "" {
		if err := godotenv.Load(cfgFilename); err == nil {
			fmt.Printf("ERROR: Failure reading ENV_FILE file: %s\n", err)
		}
		return
	}
	_, b, _, _ := runtime.Caller(0)
	cfgFilename = filepath.Join(filepath.Dir(b), "../../etc/config/config.dev.env")
	fmt.Println("CFG: ", cfgFilename)

	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Printf("ERROR: Failure reading config file: %s\n", err)
	}
}
