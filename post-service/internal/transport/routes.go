package transport

import (
	"context"
	"github.com/danielwangai/twiga-foods/post-service/internal/literals"
	"net/http"

	"github.com/danielwangai/twiga-foods/post-service/internal/svc"
	"github.com/danielwangai/twiga-foods/post-service/internal/transport/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Router ...
type Router struct {
	*mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRoutes(ctx context.Context, service svc.Svc, log *logrus.Logger, dbClient *mongo.Client) {
	e := handlers.NewEndpoints() // initialize endpoints
	r.HandleFunc(literals.HealthcheckEndpoint, e.Healthcheck(ctx, log, dbClient)).
		Methods(http.MethodGet).
		Name(literals.HealthcheckEndpointName)
	r.HandleFunc(literals.PostsBaseEndpoint, e.CreatePost(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.CreatePostEndpointName)
	r.HandleFunc(literals.PostsBaseEndpoint, e.GetPosts(ctx, service, log)).
		Methods(http.MethodGet).
		Name(literals.GetPostsEndpointName)
	r.HandleFunc(literals.PostByIDEndpoint, e.FindPostById(ctx, service, log)).
		Methods(http.MethodGet).
		Name(literals.FindPostByIDEndpointName)
}
