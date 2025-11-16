package gconf

type Server struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Protocol string `mapstructure:"protocol"`
}
