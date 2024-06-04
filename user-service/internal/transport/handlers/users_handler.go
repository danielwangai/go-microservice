package handlers

import (
	"context"
	"encoding/json"
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
	"github.com/danielwangai/twiga-foods/user-service/internal/svc"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (e *Epts) RegisterUser(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u svc.UserAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			log.WithError(literals.InvalidRegisterUserRequestPayload).Error("invalid register user payload")
			respondWithError(w, http.StatusBadRequest, literals.InvalidRegisterUserRequestPayload)
			return
		}

		svcType := convertUserApiRequestTypeToSvcType(&u)

		user, errs := service.RegisterUser(ctx, svcType)
		if errs != nil {
			respondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		log.Infof("user registered successfully: %v", user)
		respondWithJSON(w, http.StatusCreated, convertUserModelToUserAPIResponseType(user))
	}
}
