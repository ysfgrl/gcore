package gconf

type Token struct {
	Type        string `mapstructure:"type"`
	Alg         string `mapstructure:"alg"`
	Expire      string `mapstructure:"expire"`
	PrivateKey  string `mapstructure:"privateKey"`
	PublicKey   string `mapstructure:"publicKey"`
	TokenLookup string `mapstructure:"lookup"`
	AuthScheme  string `mapstructure:"scheme"`
}
