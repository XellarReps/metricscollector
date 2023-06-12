package config

import (
	"testing"
	"time"

	"github.com/XellarReps/metricscollector/internal/metrics"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testConfigPath             = "test_config/test.yaml"
	expectedAddress            = "xellar.ru:8080"
	expectedEndpoint           = "xellar.ru:8080"
	expectedTimeout            = time.Duration(10 * time.Second)
	expectedPollInterval       = time.Duration(2 * time.Second)
	expectedUpdatePerIteration = 5
)

var expectedMetrics = metrics.Collection{
	&metrics.Metric{Name: "Xellar", Type: "gauge", View: "default"},
	&metrics.Metric{Name: "BuckHashSys", Type: "gauge", View: "runtime"},
}

func TestConfig(t *testing.T) {
	cfg, err := NewConfig(testConfigPath)
	require.NoError(t, err)

	// test server
	assert.Equal(t, expectedAddress, cfg.Server.Address)

	// test agent
	assert.Equal(t, expectedEndpoint, cfg.Agent.Endpoint)
	assert.Equal(t, expectedTimeout, cfg.Agent.Client.Timeout)
	assert.Equal(t, expectedPollInterval, cfg.Agent.PollInterval)
	assert.Equal(t, expectedUpdatePerIteration, cfg.Agent.UpdatePerIteration)

	actualMetricsMap := make(map[string]metrics.Metric)
	for _, val := range cfg.Agent.Metrics {
		actualMetricsMap[val.Name] = *val
	}

	expectedMetricsMap := make(map[string]metrics.Metric)
	for _, val := range expectedMetrics {
		expectedMetricsMap[val.Name] = *val
	}

	assert.Equal(t, expectedMetricsMap, actualMetricsMap)
}
