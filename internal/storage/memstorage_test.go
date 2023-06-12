package storage

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestStorage() *MemStorage {
	storage := NewMemStorage()
	storage.gauge["xellar"] = 12345.12345
	storage.counter["abacaba"] = 12345
	return storage
}

func TestUpdateStorage(t *testing.T) {
	type params struct {
		metricType string
		name       string
		value      string
	}

	type want struct {
		gauge   map[string]float64
		counter map[string]int64
	}

	tests := []struct {
		name     string
		storage  *MemStorage
		params   params
		want     want
		checkErr bool
		err      error
	}{
		{
			name:    "update empty gauge",
			storage: NewMemStorage(),
			params: params{
				metricType: "gauge",
				name:       "lolkek",
				value:      "-123.123",
			},
			want: want{
				gauge:   map[string]float64{"lolkek": -123.123},
				counter: map[string]int64{},
			},
			checkErr: false,
		},
		{
			name:    "update empty counter",
			storage: NewMemStorage(),
			params: params{
				metricType: "counter",
				name:       "lolkek",
				value:      "123",
			},
			want: want{
				gauge:   map[string]float64{},
				counter: map[string]int64{"lolkek": 123},
			},
			checkErr: false,
		},
		{
			name:    "update gauge",
			storage: createTestStorage(),
			params: params{
				metricType: "gauge",
				name:       "xellar",
				value:      "-123.123",
			},
			want: want{
				gauge:   map[string]float64{"xellar": -123.123},
				counter: map[string]int64{"abacaba": 12345},
			},
			checkErr: false,
		},
		{
			name:    "update counter",
			storage: createTestStorage(),
			params: params{
				metricType: "counter",
				name:       "abacaba",
				value:      "-1",
			},
			want: want{
				gauge:   map[string]float64{"xellar": 12345.12345},
				counter: map[string]int64{"abacaba": 12344},
			},
			checkErr: false,
		},
		{
			name:    "error gauge value",
			storage: NewMemStorage(),
			params: params{
				metricType: "gauge",
				name:       "abacaba",
				value:      "error_type",
			},
			checkErr: true,
			err:      errors.New("invalid metric value type: strconv.ParseFloat: parsing \"error_type\": invalid syntax"),
		},
		{
			name:    "error counter value",
			storage: NewMemStorage(),
			params: params{
				metricType: "counter",
				name:       "abacaba",
				value:      "error_type",
			},
			checkErr: true,
			err:      errors.New("invalid metric value type: strconv.ParseInt: parsing \"error_type\": invalid syntax"),
		},
		{
			name:    "error metric type",
			storage: NewMemStorage(),
			params: params{
				metricType: "xellar",
				name:       "abacaba",
				value:      "123",
			},
			checkErr: true,
			err:      errors.New("invalid metric type: `xellar`"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.storage.UpdateStorage(tt.params.metricType, tt.params.name, tt.params.value)
			if tt.checkErr {
				assert.Error(t, err, tt.err)
			} else {
				require.NoError(t, err)

				assert.Equal(t, tt.want.gauge, tt.storage.gauge)
				assert.Equal(t, tt.want.counter, tt.storage.counter)
			}
		})
	}
}
