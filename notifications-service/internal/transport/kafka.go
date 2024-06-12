package transport

import (
	"github.com/Shopify/sarama"
	"github.com/danielwangai/twiga-foods/notifications-service/internal/svc"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type KafkaConsumerConfig struct {
	conn sarama.Consumer
	svc  svc.Svc
	log  *logrus.Logger
}

func NewKafkaConsumerConfig(conn sarama.Consumer, service svc.Svc, log *logrus.Logger) *KafkaConsumerConfig {
	return &KafkaConsumerConfig{
		conn, service, log,
	}
}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// ConsumeUsers listens for new users and saves to db
func (k *KafkaConsumerConfig) ConsumeNewComments(brokers []string, topic string) {
	worker, err := connectConsumer(brokers)
	if err != nil {
		panic(err)
	}

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				k.log.WithError(err)
			case msg := <-consumer.Messages():
				msgCount++
				k.log.Infof("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				k.log.Infoln("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	k.log.Infof("Processed:  %d ", msgCount)

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
