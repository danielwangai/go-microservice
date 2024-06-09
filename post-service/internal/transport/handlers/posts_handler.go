package handlers

import (
	"context"
	"encoding/json"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"github.com/danielwangai/twiga-foods/post-service/internal/svc"
	"github.com/gorilla/mux"
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

func (e *Epts) GetPosts(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := service.GetPosts(ctx)
		if err != nil {
			log.WithError(err).Error("error getting posts")
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		var res []*svc.PostAPIResponseType
		for _, post := range posts {
			res = append(res, convertPostSvcResponseTypeToAPIType(post))
		}

		log.Infof("posts retrieved successfully: %v", res)
		respondWithJSON(w, http.StatusOK, res)
	}
}

func (e *Epts) FindPostById(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from request params
		params := mux.Vars(r)

		p, err := service.FindPostById(ctx, params["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		respondWithJSON(w, http.StatusBadRequest, convertPostSvcResponseTypeToAPIType(p))
	}
}
