package config

import "github.com/XellarReps/metricscollector/internal/api"

type ServerYaml struct {
	Endpoint string `yaml:"endpoint"`
}

func serverFromYAML(yaml ServerYaml) *api.Server {
	return api.NewServer(api.ServerConfig{
		Endpoint: yaml.Endpoint,
	})
}
