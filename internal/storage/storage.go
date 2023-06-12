package storage

type Repository interface {
	UpdateStorage(metricType, name, value string) error
	GetMetricFromStorage(metricType, name string) (string, error)
	GetAllMetricsFromStorage() (map[string]float64, map[string]int64)
}
