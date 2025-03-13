package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Teamserver struct {
		Port     string `yaml:"port"`
		Password string `yaml:password`
	}
}

func LoadServerConfig(configPath string) (ServerConfig, error) {
	var config ServerConfig
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, err
}
