package config

import "github.com/XellarReps/metricscollector/internal/metrics"

type MetricYaml struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	View string `yaml:"view"`
}

func metricFromYAML(yaml MetricYaml) *metrics.Metric {
	return &metrics.Metric{
		Name: yaml.Name,
		Type: yaml.Type,
		View: yaml.View,
	}
}
