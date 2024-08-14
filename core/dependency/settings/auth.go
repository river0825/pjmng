package settings

type AuthConfig struct {
	JWTSecret string `mapstructure:"jwt_secret" validate:"required"`
}
