package storage

import (
	"fmt"
	"strconv"
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
	if metricType == "gauge" {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid metric value type: %v", err)
		}

		ms.gauge[name] = val
	} else if metricType == "counter" {
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
