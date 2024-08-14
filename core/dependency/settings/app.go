package settings

type AppConfig struct {
	Port                     string `mapstructure:"port" validate:"required"`
	AccessControlAllowOrigin string `mapstructure:"access_control_allow_origin" default:"https://service.beiqilai.com"`
}
