package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Teamserver struct {
		Port string `yaml:"port"`
	}
	Operators struct {
		Users map[string]string `yaml:"users"`
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

func GetOperators(config ServerConfig) map[string]string {
	return config.Operators.Users
}
func GetServerPort(config ServerConfig) string {
	return config.Teamserver.Port
}
