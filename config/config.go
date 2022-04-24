package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

type OpenTelemetry struct {
	URL         string `json:"url" validate:"required"`
	ServiceName string `json:"serviceName"  validate:"required"`
}

type MongoDB struct {
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBName   string `json:"dbName" validate:"required"`
}

type ServerConfig struct {
	AppVersion string `json:"appVersion" validate:"required"`
	Host       string `json:"host" validate:"required"`
	Port       string `json:"port" validate:"required"`
	IPHeader                    string `json:"IPHeader" validate:"required"`
	ShowUnknownErrorsInResponse bool   `json:"showUnknownErrorsInResponse"`
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
