package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testConfigPath   = "test_config/test.yaml"
	expectedEndpoint = "xellar.ru:8080"
)

func TestConfig(t *testing.T) {
	cfg, err := NewConfig(testConfigPath)
	require.NoError(t, err)

	assert.Equal(t, expectedEndpoint, cfg.Server.Address)
}
