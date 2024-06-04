package svc

import (
	"context"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
	"github.com/danielwangai/twiga-foods/user-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SVC struct {
	dao repo.DAO
	log *logrus.Logger
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger) Svc {
	return &SVC{dao, log}
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
	fSvc := convertUserModelToUserServiceResponseType(res)

	return fSvc, nil
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
