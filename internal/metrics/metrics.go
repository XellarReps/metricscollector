package metrics

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
)

const (
	defaultView = "default"
	runtimeView = "runtime"
)

const (
	GaugeType   = "gauge"
	CounterType = "counter"
)

const (
	randomValue = "RandomValue"
	pollCount   = "PollCount"
)

type Metric struct {
	Name string
	Type string
	View string
}

type Collection []*Metric

func (c *Collection) Collect(gauge map[string]float64, counter map[string]int64, rm *runtime.MemStats) error {
	// save runtime stats
	runtime.ReadMemStats(rm)

	runtimeStats := collectRuntimeStats(rm)

	for _, metric := range *c {
		switch metric.Type {
		case GaugeType:
			if metric.View == defaultView && metric.Name == randomValue {
				gauge[metric.Name] = float64(rand.Int()) + rand.Float64()
			} else if metric.View == defaultView && metric.Name != randomValue {
				return fmt.Errorf("unknown metric name for default view (gauge): `%s`", metric.Name)
			} else if metric.View == runtimeView {
				if val, ok := runtimeStats[metric.Name]; ok {
					gauge[metric.Name] = val
				} else {
					return fmt.Errorf("unknown metric name for runtime view (gauge): `%s`", metric.Name)
				}
			}
		case CounterType:
			if metric.View == defaultView && metric.Name == pollCount {
				counter[metric.Name] += 1
			} else if metric.View == defaultView && metric.Name != pollCount {
				return fmt.Errorf("unknown metric name for default view (counter): `%s`", metric.Name)
			} else {
				return fmt.Errorf("unknown counter view params: %s, %s, %s", metric.Name, metric.Type,
					metric.View)
			}
		default:
			return fmt.Errorf("unknown metric type: `%s`", metric.Type)
		}
	}
	return nil
}

func collectRuntimeStats(rm *runtime.MemStats) map[string]float64 {
	runtimeStats := make(map[string]float64)

	valRuntimeStats := reflect.ValueOf(rm).Elem()
	for i := 0; i < valRuntimeStats.NumField(); i++ {
		name := valRuntimeStats.Type().Field(i).Name
		value := valRuntimeStats.FieldByName(name).Interface()

		switch val := value.(type) {
		case uint32:
			runtimeStats[name] = float64(val)
		case uint64:
			runtimeStats[name] = float64(val)
		case float64:
			runtimeStats[name] = val
		}
	}

	return runtimeStats
}
