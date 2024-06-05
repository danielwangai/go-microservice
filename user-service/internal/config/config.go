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
	envPrefix = "SERVICE_USER"
)

// WebServerConfig ...
type WebServerConfig struct {
	Port string `envconfig:"SERVICE_USER_SERVER_PORT" split_words:"true"`
}

type MongoConfig struct {
	DbURL  string `envconfig:"SERVICE_USER_DATABASE_URL"`
	DbName string `envconfig:"SERVICE_USER_DATABASE_NAME"`
}

type KafkaConfig struct {
	BootstrapServers string `envconfig:"SERVICE_USER_KAFKA_BOOTSTRAP_SERVERS" split_words:"true"`
	ClientID         string `envconfig:"SERVICE_USER_KAFKA_CLIENT_ID"`
	AutoOffsetReset  string `envconfig:"SERVICE_USER_KAFKA_AUTO_OFFSET_RESET"`
	Topic            string `envconfig:"SERVICE_USER_KAFKA_CREATE_USER_TOPIC"`
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
