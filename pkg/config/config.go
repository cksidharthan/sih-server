package config

import (
	"fmt"
	envParser "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"9090"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func New() (*Config, error) {
	var info Config
	// get project root path (to load .env file)
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// check if .env file exists in project root path (if yes, load it | if no, skip and use default values)
	if _, err := os.Stat(rootPath + "/config.env"); err == nil {
		// load .env file
		logrus.Infof("Loading .env file from %s", rootPath)
		err = godotenv.Load(rootPath + "/config.env")
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	// Parse the config file and initialize the env object with required values.
	if err := envParser.Parse(&info); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("URL: http://localhost:%s\n", info.Port)
	fmt.Println("-----------------------------------------")

	return &info, nil
}
