package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/XellarReps/metricscollector/internal/agent"
	"github.com/XellarReps/metricscollector/internal/api"
	"github.com/XellarReps/metricscollector/internal/metrics"
)

type Config struct {
	Server *api.Server
	Agent  *agent.Agent
}

type ConfigYaml struct {
	Server  ServerYaml   `yaml:"server"`
	Agent   AgentYaml    `yaml:"agent"`
	Metrics []MetricYaml `yaml:"metrics"`
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

func configFromYAML(yaml ConfigYaml) (*Config, error) {
	server := serverFromYAML(yaml.Server)

	var ms metrics.Collection
	for _, val := range yaml.Metrics {
		metric := metricFromYAML(val)
		ms = append(ms, metric)
	}

	agentClient := agentFromYAML(yaml.Agent, ms)

	return &Config{
		Server: server,
		Agent:  agentClient,
	}, nil
}
