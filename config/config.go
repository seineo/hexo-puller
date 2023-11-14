package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var CONFIG_PATH = "config/config.yaml"
var Conf *Config

type Config struct {
	Url  string `yaml:"url"`
	Path string `yaml:"path"`
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
