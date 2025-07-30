package env_config

import (
	"github.com/caarlos0/env/v11"
)

type (
	httpServerConfig struct {
		cfg httpServerEnvConfig
	}

	httpServerEnvConfig struct {
		Port int64 `env:"HTTP_PORT,required"`
	}
)

// NewHttpServerConfig -
func NewHttpServerConfig() (*httpServerConfig, error) {
	cfg := httpServerEnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &httpServerConfig{
		cfg: cfg,
	}, nil
}

// Port -
func (c *httpServerConfig) Port() int64 {
	return c.cfg.Port
}
