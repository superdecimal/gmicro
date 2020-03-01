package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration ...
type Configuration struct {
	Port int `default:"3000"`
}

const prefix = "GM"

func Read() (
	*Configuration,
	error,
) {
	var configuration Configuration

	if err := envconfig.Process(prefix, &configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}
