package config

import (
	"os"

	"github.com/soerjadi/stockist/internal/pkg/util"
	"gopkg.in/gcfg.v1"
)

var configFilePaths = map[string]string{
	"PRODUCTION":  "/etc/stockist/config.ini",
	"DEVELOPMENT": "../../files/config.ini",
}

func Init() (*Config, error) {
	cfg = &Config{}

	configFilePath := configFilePaths[util.GetENV()]

	config, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = gcfg.ReadStringInto(cfg, string(config))
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// GetConfig returns config object
func GetConfig() *Config {
	return cfg
}
