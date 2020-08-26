package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

// ReadEnvFromYaml ...
// Creates environmental variables from config.yml
func ReadEnvFromYaml() <-chan Config {
	task := make(chan Config)
	go func() {
		var cfg Config
		FetchYaml("config.yml", &cfg)
		GrabEnv(&cfg)
		task <- cfg
	}()
	return task
}

// Config ...
type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD`
	} `yaml:"database"`
}

// FetchYaml ...
// Fetches data to "c *Config" from "config.yml"
func FetchYaml(path string, c *Config) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

// GrabEnv ...
// Grabs env data and stores it to "c *Config"
func GrabEnv(c *Config) {
	err := envconfig.Process("", c)
	if err != nil {
		panic(err)
	}
}
