package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config struct {
	Auth struct {
		SessionKey string `yaml:"session_key"`
		LogoutKey  string `yaml:"logout_key"`
		Working    bool   `yaml:"working"`
	}
	Database struct {
		Dsn      string `yaml:"dsn"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Ssl      string `yaml:"ssl"`
	}
	App struct {
		HttpPort string `yaml:"http_port"`
		GrpcPort string `yaml:"grpc_port"`
	}
	GRPC struct {
		UserAPI string `yaml:"user_api"`
	}
}

var Config = config{}

func ReadConf(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	c := &Config
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", filename, err)
	}
	return nil
}
