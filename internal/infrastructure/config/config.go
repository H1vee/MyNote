package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		User     string `yaml:"user"`
		PassWord string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	JWT struct {
		Secret                   string `yaml:"secret"`
		AccesTokenTTLMinute      uint   `yaml:"acces_token_ttl_minute:"`
		refresh_token_ttl_minute uint   `yaml:"refresh_token_ttl_minute:"`
	}
}

func Load(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding yaml: %v", err)
	}
	return &cfg
}
