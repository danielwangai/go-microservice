package svc

import (
	"github.com/danielwangai/twiga-foods/user-service/internal/repo/mongo"
	"github.com/sirupsen/logrus"
)

type SVC struct {
	dao repo.DAO
	log *logrus.Logger
}

// New returns a new Svc object
func New(dao repo.DAO, log *logrus.Logger) Svc {
	return &SVC{dao, log}
}
