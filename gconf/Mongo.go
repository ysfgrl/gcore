package gconf

type Mongo struct {
	Url      string `mapstructure:"url"`
	Db       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}
