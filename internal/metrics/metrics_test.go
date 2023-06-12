package metrics

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var expectedNamesCollectGauge = []string{"Alloc", "RandomValue"}
var expectedNamesCollectCounter = []string{"PollCount"}

func TestCollect(t *testing.T) {
	rm := runtime.MemStats{}

	type params struct {
		gauge   map[string]float64
		counter map[string]int64
		rm      *runtime.MemStats
	}

	collection := Collection{
		&Metric{Name: "RandomValue", Type: "gauge", View: "default"},
		&Metric{Name: "PollCount", Type: "counter", View: "default"},
		&Metric{Name: "Alloc", Type: "gauge", View: "runtime"},
	}

	tests := []struct {
		name             string
		collection       Collection
		params           params
		wantNamesGauge   []string
		wantNamesCounter []string
		wantCount        int64
	}{
		{
			name:       "empty maps",
			collection: collection,
			params: params{
				gauge:   make(map[string]float64),
				counter: make(map[string]int64),
				rm:      &rm,
			},
			wantNamesGauge:   expectedNamesCollectGauge,
			wantNamesCounter: expectedNamesCollectCounter,
			wantCount:        1,
		},
		{
			name:       "not empty maps",
			collection: collection,
			params: params{
				gauge: map[string]float64{
					"RandomValue": 123.123,
				},
				counter: map[string]int64{
					"PollCount": 123,
				},
				rm: &rm,
			},
			wantNamesGauge:   expectedNamesCollectGauge,
			wantNamesCounter: expectedNamesCollectCounter,
			wantCount:        124,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.collection.Collect(tt.params.gauge, tt.params.counter, tt.params.rm)
			require.NoError(t, err)

			var actualNamesGauge []string
			for key := range tt.params.gauge {
				actualNamesGauge = append(actualNamesGauge, key)
			}

			assert.ElementsMatch(t, actualNamesGauge, tt.wantNamesGauge)

			var actualNamesCounter []string
			for key := range tt.params.counter {
				actualNamesCounter = append(actualNamesCounter, key)
			}

			assert.ElementsMatch(t, actualNamesCounter, tt.wantNamesCounter)

			assert.Equal(t, tt.params.counter["PollCount"], tt.wantCount)
		})
	}
}

var expectedNamesRuntimeStats = []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc",
	"HeapIdle", "HeapInuse", "HeapObjects", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys",
	"MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs",
	"StackInuse", "StackSys", "Sys", "TotalAlloc",
}

func TestCollectRuntimeStats(t *testing.T) {
	rm := runtime.MemStats{}

	runtime.ReadMemStats(&rm)

	rmStats := collectRuntimeStats(&rm)
	var actualNames []string
	for key := range rmStats {
		actualNames = append(actualNames, key)
	}

	assert.ElementsMatch(t, expectedNamesRuntimeStats, actualNames)
}
