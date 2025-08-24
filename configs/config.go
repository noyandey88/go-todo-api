package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port         int `yaml:"port"`
		ReadTimeout  int `yaml:"readTimeout"`
		WriteTimeout int `yaml:"writeTimeout"`
	} `yaml:"server"`

	Database struct {
		Host               string `yaml:"host"`
		Port               int    `yaml:"port"`
		User               string `yaml:"user"`
		Password           string `yaml:"password"`
		Name               string `yaml:"name"`
		SSLMode            string `yaml:"sslmode"`
		TimeZone           string `yaml:"timezone"`
		SuperAdmin         string `yaml:"superAdmin"`
		SuperAdminPassword string `yaml:"superAdminPassword"`
	} `yaml:"database"`

	JWT struct {
		Secret    string `yaml:"secret"`
		ExpiresIn int    `yaml:"expiresIn"`
	} `yaml:"jwt"`

	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

var AppConfig *Config

func LoadConfig() *Config {
	file, err := os.Open("configs/config-prod.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	AppConfig = &Config{}
	err = decoder.Decode(AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	return AppConfig
}
