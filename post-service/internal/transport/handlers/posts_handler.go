package handlers

import (
	"context"
	"encoding/json"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"github.com/danielwangai/twiga-foods/post-service/internal/svc"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (e *Epts) CreatePost(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p svc.PostAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&p); err != nil {
			log.WithError(literals.InvalidCreatePostRequestPayload).Error("invalid create post payload")
			respondWithError(w, http.StatusBadRequest, literals.InvalidCreatePostRequestPayload)
			return
		}

		svcType := convertCreatePostApiRequestTypeToSvcType(&p)

		post, errs := service.CreatePost(ctx, svcType)
		if errs != nil {
			respondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		log.Infof("post created successfully: %v", post)
		respondWithJSON(w, http.StatusCreated, convertPostSvcResponseTypeToAPIType(post))
	}
}
