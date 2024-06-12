package literals

const (
	TimestampFormat         = "2006-01-02 15:04:05"
	RoutePrefix             = "post-service/v1"
	DatabaseName            = "notifications_service_db"
	HealthcheckEndpoint     = "/healthcheck"
	HealthcheckEndpointName = "healthcheck"

	// collections
	PostsCollection      = "posts"
	UsersCollection      = "users"
	UserFollowCollection = "user_followers"
	CommentsCollection   = "comments"

	// kafka topics
	NewUserTopic    = "new-users-events"
	NewPostTopic    = "new-post-notification-events"
	NewCommentTopic = "new-comment-notification-events"
	FollowUserTopic = "follow-user-notification-events"
)
