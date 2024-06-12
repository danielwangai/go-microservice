package transport

import (
	"context"
	"encoding/json"
	"fmt"
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

// ConsumeNewComments listens for new comments and saves to db
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

				var c *svc.CommentServiceRequestType
				if err = json.Unmarshal(msg.Value, &c); err != nil {
					k.log.WithError(err).Error("Error unmarshalling comment data from kafka")
					break
				}
				// save to database
				comment, errs := k.svc.AddComment(context.Background(), c)
				if errs != nil {
					k.log.Errorf("error adding comment in kafka consumer: %v", errs)
				}

				k.log.Infof("comment inserted to database after consumption from kafka: %v", comment)
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

// ConsumeNewPost listens for new posts and saves to db
func (k *KafkaConsumerConfig) ConsumeNewPost(brokers []string, topic string) {
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

				var p *svc.PostServiceRequestType
				if err = json.Unmarshal(msg.Value, &p); err != nil {
					k.log.WithError(err).Error("Error unmarshalling post data from kafka")
					break
				}
				// save to database
				comment, errs := k.svc.CreatePost(context.Background(), p)
				if errs != nil {
					k.log.Errorf("error adding post in kafka consumer: %v", errs)
				}

				k.log.Infof("post inserted to database after consumption from kafka: %v", comment)
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

// ConsumeUserFollowInfo listens for new user followings and saves to db
func (k *KafkaConsumerConfig) ConsumeUserFollowInfo(brokers []string, topic string) {
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

				var f *svc.UserFollowerServiceResponseType
				if err = json.Unmarshal(msg.Value, &f); err != nil {
					k.log.WithError(err).Error("Error unmarshalling post data from kafka")
					break
				}
				// save to database
				comment, errs := k.svc.StoreFollowInfo(context.Background(), f.Follower.ID, f.Followed.ID)
				if errs != nil {
					k.log.Errorf("error adding follow info in kafka consumer: %v", errs)
				}

				k.log.Infof("post inserted to database after consumption from kafka: %v", comment)
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

// ConsumeUsers listens for new users and saves to db
func (k *KafkaConsumerConfig) ConsumeUsers(brokers []string, topic string) {
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
	fmt.Println("Consumer started ")
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
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))

				var u *svc.UserServiceRequestType
				if err = json.Unmarshal(msg.Value, &u); err != nil {
					k.log.WithError(err).Error("Error unmarshalling user data from kafka")
					break
				}
				// save to database
				user, err := k.svc.AddUser(context.Background(), u)
				if err != nil {
					k.log.WithError(err).Errorf("error adding user in kafka consumer: %v", user)
					//return literals.UserInsertionError
				}

				k.log.Infof("user inserted to database after consumption from kafka: %v", user)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
