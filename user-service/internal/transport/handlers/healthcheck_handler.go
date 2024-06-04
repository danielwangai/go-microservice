package handlers

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *Epts) Healthcheck(ctx context.Context, log *logrus.Logger, dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		res := make(map[string]string)
		if err := dbClient.Ping(context.TODO(), nil); err != nil {
			res["db"] = "db connection failed"
			res["overallSystemStatus"] = "not okay"
			respondWithJSON(w, http.StatusInternalServerError, res)
			return
		}

		res["overallSystemStatus"] = "ok"
		res["db"] = "db connection successful"

		respondWithJSON(w, http.StatusOK, res)
	}
}
