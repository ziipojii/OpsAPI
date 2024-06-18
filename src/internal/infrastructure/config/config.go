package config

import (
    "github.com/kelseyhightower/envconfig"
    "github.com/joho/godotenv"
)

type Configuration struct {

	HttpServerPort int    `envconfig:"HTTP_SERVER_PORT" default:"8888"`
    ServiceName    string `envconfig:"SERVICE_NAME" default:"tix-github.com/ziipojii/OpsAPI.git"`
	AppEnv         string `envconfig:"APP_ENV" default:"production"`

}


func Config() (*Configuration, error) {

	// Load environment variables 
	err := godotenv.Load()

	var configuration Configuration
	err = envconfig.Process("", &configuration)
	if err != nil {
		return nil, err
	}


	return &configuration, nil
}
