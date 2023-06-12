package config

import (
	"time"

	"github.com/XellarReps/metricscollector/internal/agent"
	"github.com/XellarReps/metricscollector/internal/metrics"
)

type AgentYaml struct {
	Endpoint           string        `yaml:"endpoint"`
	Timeout            time.Duration `yaml:"timeout"`
	PollInterval       time.Duration `yaml:"poll_interval"`
	UpdatePerIteration int           `yaml:"update_per_iteration"`
}

func agentFromYAML(yaml AgentYaml, metrics metrics.Collection) *agent.Agent {
	return agent.NewAgent(agent.Config{
		Endpoint:           yaml.Endpoint,
		Timeout:            yaml.Timeout,
		PollInterval:       yaml.PollInterval,
		UpdatePerIteration: yaml.UpdatePerIteration,
		Metrics:            metrics,
	})
}
