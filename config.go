package main

import (
	"github.com/kelseyhightower/envconfig"
)

type appConfig struct {
	Port          int    `required:"true"`
	Env           string `default:"local"`
	ComponentName string `default:"no-name"`
}

func (c *appConfig) IsLocal() bool {
	return c.Env == "local"
}

func loadAppConfig() (*appConfig, error) {
	var config appConfig
	err := envconfig.Process("", &config)
	return &config, err
}
