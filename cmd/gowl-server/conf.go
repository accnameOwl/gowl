package main

import (
	"os"

	"github.com/gofiber/fiber"

	//! To be implemented
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// ReadEnvFromYaml ...
// Creates environmental variables from config.yml
// ! wrong way to use channels ???
func ReadEnvFromYaml() <-chan Config {
	task := make(chan Config)
	go func() {
		var cfg Config
		FetchYaml("config.yml", &cfg)
		//!	GrabEnv(&cfg)
		task <- cfg
	}()
	return task
}

// Config ...
type Config struct {
	PPROF struct {
		Active  bool
		CPU     bool
		Mem     bool
		MemRate int
		Trace   bool
	}
	Server struct {
		Port string `yaml:"port"` //, envconfig:"SERVER_PORT"`
		Host string `yaml:"host"` //, envconfig:"SERVER_HOST`
	} `yaml:"server"`
	Fiber struct { // ! envconfig values
		ServerHeader              string `yaml:"server_header"`
		StrictRouting             bool   `yaml:"strict_routing"`
		CaseSensitive             bool   `yaml:"case_sensitive"`
		Immutable                 bool   `yaml:"immutable"`
		UnescapePath              bool   `yaml:"unescape_path"`
		ETag                      bool   `yaml:"etag"`
		Prefork                   bool   `yaml:"prefork"`
		BodyLimit                 int    `yaml:"body_limit"`
		Concurrency               int    `yaml:"concurrency"`
		DisableHeaderNormalizing  bool   `yaml:"disable_header_normalizing"`
		DisableKeepalive          bool   `yaml:"disable_keep_alive"`
		DisableDefaultDate        bool   `yaml:"disable_default_date"`
		DisableDefaultContentType bool   `yaml:"disable_default_content_type"`
		DisableStartupMessage     bool   `yaml:"disable_startup_message"`
		ReadBufferSize            int    `yaml:"read_buffer_size"`
		WriteBufferSize           int    `yaml:"write_buffer_size"`
		CompressedFileSuffix      string `yaml:"compressed_file_suffix"`
	} `yaml:"fiber"`
	Database struct {
		Username string `yaml:"user"` //, envconfig:"DB_USERNAME"`
		Password string `yaml:"pass"` //, envconfig:"DB_PASSWORD`
	} `yaml:"database"`
}

// FetchFiberSettings ...
// Returns a new Settings type with settings from config.yml
// ! Wrong way to use channels ??
func FetchFiberSettings(c *Config) <-chan fiber.Settings {
	out := make(chan fiber.Settings)
	go func() {
		out <- fiber.Settings{
			ServerHeader:              c.Fiber.ServerHeader,
			StrictRouting:             c.Fiber.StrictRouting,
			CaseSensitive:             c.Fiber.CaseSensitive,
			Immutable:                 c.Fiber.Immutable,
			UnescapePath:              c.Fiber.UnescapePath,
			ETag:                      c.Fiber.ETag,
			Prefork:                   c.Fiber.Prefork,
			BodyLimit:                 c.Fiber.BodyLimit,
			Concurrency:               c.Fiber.Concurrency,
			DisableKeepalive:          c.Fiber.DisableKeepalive,
			DisableDefaultDate:        c.Fiber.DisableDefaultDate,
			DisableDefaultContentType: c.Fiber.DisableDefaultContentType,
			DisableHeaderNormalizing:  c.Fiber.DisableHeaderNormalizing,
			DisableStartupMessage:     c.Fiber.DisableStartupMessage,
			ReadBufferSize:            c.Fiber.ReadBufferSize,
			WriteBufferSize:           c.Fiber.WriteBufferSize,
			CompressedFileSuffix:      c.Fiber.CompressedFileSuffix,
		}
		close(out)
	}()
	return out
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
