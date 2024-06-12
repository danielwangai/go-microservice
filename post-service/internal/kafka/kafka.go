package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	log      *logrus.Logger
	brokers  []string
	topics   map[string]string
}

func ConnectProducer(topics map[string]string, brokers []string, log *logrus.Logger) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{conn, log, brokers, topics}, nil
}

func (k *KafkaProducer) PushMessageToQueue(topic, key string, message []byte) error {
	defer k.producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: k.topics[topic],
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)

	return nil
}
