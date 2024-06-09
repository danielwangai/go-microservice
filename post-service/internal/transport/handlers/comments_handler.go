package handlers

import (
	"context"
	"encoding/json"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"github.com/danielwangai/twiga-foods/post-service/internal/svc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (e *Epts) AddComment(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var c svc.CommentAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			log.WithError(literals.InvalidAddCommentRequestPayload).Error("invalid add comment payload")
			respondWithError(w, http.StatusBadRequest, literals.InvalidAddCommentRequestPayload)
			return
		}

		// postId

		svcType := convertAddCommentAPIRequestTypeToSvcType(&c)
		svcType.PostID = params["id"]

		comment, errs := service.AddComment(ctx, svcType)
		if errs != nil {
			respondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		log.Infof("post created successfully: %v", comment)
		respondWithJSON(w, http.StatusCreated, convertCommentSvcResponseTypeToAPIType(comment))
	}
}
