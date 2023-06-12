package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	templatePath = "templates/url.tpl"
)

func TestCreateMetricURL(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]any
		want   string
	}{
		{
			name: "success test 1",
			params: map[string]any{
				"hostname": "abacaba.com:8080",
				"type":     "gauge",
				"name":     "testik",
				"value":    "123.123",
			},
			want: "http://abacaba.com:8080/update/gauge/testik/123.123",
		},
		{
			name: "success test 2",
			params: map[string]any{
				"hostname": "abacaba.com:8080",
				"type":     "counter",
				"name":     "testik",
				"value":    "123",
			},
			want: "http://abacaba.com:8080/update/counter/testik/123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := CreateMetricURL(templatePath, tt.params)
			require.NoError(t, err)

			assert.Equal(t, tt.want, actual)
		})
	}
}
