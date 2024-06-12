package handlers

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -package mocks -destination=./mocks/mock_endpoints.go . Endpoints
type Endpoints interface {
	Healthcheck(ctx context.Context, log *logrus.Logger, dbClient *mongo.Client) http.HandlerFunc
}

type Epts struct{}

func NewEndpoints() Endpoints {
	return &Epts{}
}
