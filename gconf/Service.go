package gconf

type Service struct {
	Url       string `mapstructure:"url"`
	Token     string `mapstructure:"token"`
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
}
