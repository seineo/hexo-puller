package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var CONFIG_PATH = "config/config.yaml"
var Conf *Config

type TLS struct {
	Crt string `yaml:"crt"`
	Key string `yaml:"key"`
}

type Config struct {
	Tls TLS
}

func newConfig(path string) *Config {
	var conf Config
	configContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %s\n", path)
	}
	err = yaml.Unmarshal(configContent, &conf)
	if err != nil {
		log.Fatalf("failed to load config, err: %s\n", err.Error())
	}
	return &conf
}

func GetConfig() *Config {
	if Conf == nil {
		Conf = newConfig(CONFIG_PATH)
	}
	return Conf
}
