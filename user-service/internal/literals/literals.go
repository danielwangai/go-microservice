package literals

const (
	RoutePrefix             = "user-service/v1"
	DatabaseName            = "user_service_db"
	HealthcheckEndpoint     = "/healthcheck"
	HealthcheckEndpointName = "healthcheck"

	// collections
	UsersCollection = "users"

	// user endpoints
	UsersBaseEndpoint        = "/users"
	UserRegisterEndpoint     = "/users/register"
	UsersLoginEndpoint       = "/users/login"
	RegisterUserEndpointName = "register-user-endpoint-name"
	LoginEndpointName        = "login-endpoint-name"
)
