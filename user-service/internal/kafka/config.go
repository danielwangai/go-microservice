package kafka

import (
	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	producer  sarama.SyncProducer
	serverUrl string
	topic     string
}

func New(producer sarama.SyncProducer, serverUrl string, topic string) *KafkaConfig {
	return &KafkaConfig{
		producer:  producer,
		serverUrl: serverUrl,
		topic:     topic,
	}
}
