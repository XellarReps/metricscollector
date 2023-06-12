package storage

type Repository interface {
	UpdateStorage(metricType, name, value string) error
}
