package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

const defaultConfigPath = "development-config.yaml"

type Config struct {
	Service Service `yaml:"service"`
}

type Service struct {
	LogLevel string `yaml:"logLevel" env:"LOG_LEVEL" env-default:"info"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) ReadConfig() error {
	configPath := os.Getenv("CONFIG_PATH")

	configExists, err := fileExists(configPath)
	if err != nil {
		return err
	}
	if !configExists {
		configPath = defaultConfigPath
	}

	slog.Info(fmt.Sprintf("Loading config from %s", configPath))
	err = cleanenv.ReadConfig(configPath, c)
	if err != nil {
		return errors.Wrapf(err, "failed to read config from %s", configPath)
	}
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
