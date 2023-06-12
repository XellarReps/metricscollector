package storage

import (
	"fmt"
	"strconv"
)

const (
	gauge   = "gauge"
	counter = "counter"
)

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (ms *MemStorage) UpdateStorage(metricType, name, value string) error {
	if metricType == gauge {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid metric value type: %v", err)
		}

		ms.gauge[name] = val
	} else if metricType == counter {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid metric value type: %v", err)
		}

		ms.counter[name] += val
	} else {
		return fmt.Errorf("invalid metric type `%s`", metricType)
	}

	return nil
}

func (ms *MemStorage) GetMetricFromStorage(metricType, name string) (string, error) {
	if metricType == gauge {
		if val, ok := ms.gauge[name]; ok {
			return fmt.Sprintf("%f", val), nil
		} else {
			return "", fmt.Errorf("metric with name `%s` not found", name)
		}
	} else if metricType == counter {
		if val, ok := ms.counter[name]; ok {
			return fmt.Sprintf("%d", val), nil
		} else {
			return "", fmt.Errorf("metric with name `%s` not found", name)
		}
	}

	return "", nil
}

func (ms *MemStorage) GetAllMetricsFromStorage() (map[string]float64, map[string]int64) {
	return ms.gauge, ms.counter
}
