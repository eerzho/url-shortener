package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	// from flag
	log.Print("Starting to get --config flag")
	path, err := fetchPathFlag()
	if err == nil {
		return MustLoadByPath(path)
	}
	log.Print(err.Error())

	// from env
	log.Print("Starting to get CONFIG_PATH variable")
	path, err = fetchPathEnv()
	if err == nil {
		return MustLoadByPath(path)
	}
	log.Print(err.Error())

	panic("config path is empty")
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &config
}

func fetchPathFlag() (string, error) {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		return "", fmt.Errorf("--config flag is not set")
	}

	return path, nil
}

func fetchPathEnv() (string, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return "", fmt.Errorf("CONFIG_PATH variable is not set")
	}

	return path, nil
}
