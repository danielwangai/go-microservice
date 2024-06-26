package svc

import (
	"context"
	"encoding/json"
	k "github.com/danielwangai/twiga-foods/user-service/internal/kafka"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
	"github.com/danielwangai/twiga-foods/user-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SVC struct {
	dao   repo.DAO
	log   *logrus.Logger
	kafka *k.KafkaProducer
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger, kafka *k.KafkaProducer) Svc {
	return &SVC{dao, log, kafka}
}

func (s *SVC) RegisterUser(ctx context.Context, u *UserServiceRequestType) (*UserServiceResponseType, literals.Error) {
	errs := validateRegisterUserInputs(u)
	if len(errs) > 0 {
		return nil, errs
	}

	// hash password
	hash, err := hashPassword([]byte(u.Password))
	if err != nil {
		errs["password"] = PasswordHashingError.Error()
		s.log.WithError(err).Error(PasswordHashingError)
		return nil, errs
	}

	u.Password = hash

	// convert user service type to model layer type
	uModel := convertUserServiceRequestTypeToModelType(u)

	// save to db
	res, errs := s.dao.RegisterUser(ctx, uModel)
	if len(errs) > 0 {
		return nil, errs
	}

	// convert from user model type to user response type for service layer
	uSvc := convertUserModelToUserServiceResponseType(res)

	// send to kafka
	uByte, err := json.Marshal(uSvc)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}
	err = s.kafka.PushMessageToQueue(literals.NewUserTopic, uSvc.ID, uByte)
	if err != nil {
		s.log.WithError(err).Errorf("failed to write user: %v to kafka topic", uSvc)
		return nil, map[string]string{"error": err.Error()}
	}

	return uSvc, nil
}

func (s *SVC) FollowUser(ctx context.Context, id1, id2 string) (*UserFollowerServiceResponseType, error) {
	follower, err := s.dao.FindUserByID(ctx, id1)
	if err != nil {
		s.log.WithError(err).Errorf("invalid id for follower user: %s", id1)
		return nil, err
	}

	// check if user to be followed exists
	followed, err := s.dao.FindUserByID(ctx, id2)
	if err != nil {
		s.log.WithError(err).Errorf("invalid id for followed user: %s", id2)
		return nil, err
	}

	// follow user
	followObj, err := s.dao.FollowUser(ctx, follower, followed)
	if err != nil {
		return nil, err
	}

	followSvc := convertUserFollowModelToServiceResponseType(followObj)

	// send to kafka
	followByte, err := json.Marshal(followSvc)
	if err != nil {
		return nil, literals.ObjectToByteArrayConversionError
	}
	err = s.kafka.PushMessageToQueue(literals.FollowUserTopic, followSvc.ID, followByte)
	if err != nil {
		s.log.WithError(err).Errorf("failed to write follow record: %v to kafka topic", followByte)
		return nil, literals.FailedToPublishMessageToKafka
	}

	return followSvc, nil
}

func (s *SVC) FindUserById(ctx context.Context, id string) (*UserServiceResponseType, error) {
	user, err := s.dao.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertUserModelToUserServiceResponseType(user), nil
}

func (s *SVC) FindUserByEmail(ctx context.Context, email string) (*UserServiceResponseType, error) {
	user, err := s.dao.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return convertUserModelToUserServiceResponseType(user), nil
}

func hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
