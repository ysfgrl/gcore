package gconf

import (
	"github.com/spf13/viper"
)

//var readFactory sync.Once
//var isInit = false

func FromYaml(file string, conf interface{}) error {
	return read("yaml", file, conf)
}
func FromJson(file string, conf interface{}) error {
	return read("json", file, conf)
}
func FromEnv(file string, conf interface{}) error {
	return read("env", file, conf)
}

func read(typ string, file string, conf interface{}) error {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.SetConfigFile(file)
	v.SetConfigType(typ)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(conf)
	if err != nil {
		return err
	}
	return nil
}

type Conf struct {
	Mongo    *Mongo              `mapstructure:"mongo"`
	Token    *Token              `mapstructure:"token"`
	Server   *Server             `mapstructure:"server"`
	Storage  *Storage            `mapstructure:"storage"`
	Services map[string]*Service `mapstructure:"services"`
}

func (c *Conf) GetService(service string) *Service {
	if c.Services == nil {
		return &Service{}
	}
	value, exists := c.Services[service]
	if exists {
		return value
	} else {
		return &Service{}
	}
}
