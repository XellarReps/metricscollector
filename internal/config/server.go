package config

import "github.com/XellarReps/metricscollector/internal/api"

type ServerYaml struct {
	Address string `yaml:"address"`
}

func serverFromYAML(yaml ServerYaml) *api.Server {
	return api.NewServer(api.ServerConfig{
		Address: yaml.Address,
	})
}
