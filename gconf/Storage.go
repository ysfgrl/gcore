package gconf

type Storage struct {
	Url       string `mapstructure:"url"`
	Port      int    `mapstructure:"port"`
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Region    string `mapstructure:"region"`
	Secure    bool   `mapstructure:"secure"`
}
