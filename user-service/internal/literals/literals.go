package literals

const (
	RoutePrefix             = "user-service/v1"
	DatabaseName            = "user_service_db"
	HealthcheckEndpoint     = "/healthcheck"
	HealthcheckEndpointName = "healthcheck"

	// collections
	UsersCollection      = "users"
	UserFollowCollection = "user_followers"

	// user endpoints
	UsersBaseEndpoint        = "/users"
	UserRegisterEndpoint     = "/users/register"
	UsersLoginEndpoint       = "/users/login"
	RegisterUserEndpointName = "register-user-endpoint-name"
	LoginEndpointName        = "login-endpoint-name"
	FollowUserEndpoint       = "/follow-user"
	FollowUserEndpointName   = "follow-user-endpoint-name"

	// kafka topics
	NewUserTopic    = "new-users-events"
	NewPostTopic    = "new-post-notification-events"
	NewCommentTopic = "new-comment-notification-events"
	FollowUserTopic = "follow-user-notification-events"
)
