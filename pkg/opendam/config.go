package opendam

import (
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Machinery      *config.Config
	WorkingDir     string `envconfig:"WORKING_DIR" default:"/dir"`
	OpenDAMHost    string `envconfig:"OPENDAM_HOST"`
	BlobConnection string `envconfig:"BLOB_CONNECTION"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}

	cfg.Machinery, err = config.NewFromEnvironment(true)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
