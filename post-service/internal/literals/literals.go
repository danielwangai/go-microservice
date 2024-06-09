package literals

const (
	TimestampFormat         = "2006-01-02 15:04:05"
	RoutePrefix             = "post-service/v1"
	DatabaseName            = "post_service_db"
	HealthcheckEndpoint     = "/healthcheck"
	HealthcheckEndpointName = "healthcheck"

	// collections
	PostsCollection    = "posts"
	UsersCollection    = "users"
	CommentsCollection = "comments"

	// post endpoints
	PostsBaseEndpoint            = "/posts"
	PostByIDEndpoint             = PostsBaseEndpoint + "/{id}"
	PostCommentBaseEndpoint      = PostsBaseEndpoint + "/{id}/comments"
	CreatePostEndpointName       = "create-post-endpoint"
	GetPostsEndpointName         = "get-posts-endpoint"
	FindPostByIDEndpointName     = "get-posts-endpoint"
	AddCommentToPostEndpointName = "post-comment-endpoint"
)
