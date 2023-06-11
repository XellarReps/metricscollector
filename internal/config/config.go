package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/XellarReps/metricscollector/internal/api"
)

type Config struct {
	Server *api.Server
}

type ConfigYaml struct {
	Server ServerYaml `yaml:"server"`
}

func NewConfig(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg ConfigYaml
	if err := yaml.UnmarshalStrict(buf, &cfg); err != nil {
		return nil, err
	}

	return configFromYAML(cfg)
}

func configFromYAML(cfg ConfigYaml) (*Config, error) {
	server := serverFromYAML(cfg.Server)
	return &Config{
		Server: server,
	}, nil
}
