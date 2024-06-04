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

func (e *Epts) Login(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u svc.UserLoginAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			log.WithError(err).Error("invalid request body")
			respondWithError(w, http.StatusBadRequest, literals.InvalidLoginRequestPayload)
			return
		}

		// find user by email
		user, err := service.FindUserByEmail(ctx, u.Email)
		if err != nil {
			log.WithError(err).Errorf("login failed due to invalid credentials from email: %s", u.Email)
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		// verify password
		if !svc.CheckPasswordHash(u.Password, user.Password) {
			log.WithError(err).Errorf("login failed due to invalid password")
			respondWithError(w, http.StatusBadRequest, literals.InvalidLoginCredentials)
			return
		}

		// generate jwt token
		tokenString, err := svc.GenerateJWT(user)
		if err != nil {
			log.WithError(err).Errorf("login failed due to invalid JWT claims for email: %s", u.Email)
			respondWithError(w, http.StatusBadRequest, literals.LoginAttemptFailed)
			return
		}

		res := map[string]string{
			"message": "login successful",
			"token":   tokenString,
		}

		log.Infof("user logged in successfully with email: %s", u.Email)
		respondWithJSON(w, http.StatusOK, res)
		return
	}
}
