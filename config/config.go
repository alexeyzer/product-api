package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config struct {
	Auth struct {
		UserInfoKey string `yaml:"user_info_key"`
		SessionKey  string `yaml:"session_key"`
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
		UserAPI      string `yaml:"user_api"`
		RecognizeAPI string `yaml:"recognize_api"`
	}
	S3 struct {
		BucketName string `yaml:"bucket_name"`
		ID         string `yaml:"id"`
		Key        string `yaml:"key"`
		Endpoint   string `yaml:"endpoint"`
		Region     string `yaml:"region"`
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
