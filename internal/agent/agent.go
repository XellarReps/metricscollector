package agent

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/XellarReps/metricscollector/internal/metrics"
	"github.com/XellarReps/metricscollector/internal/utils"
	"github.com/go-resty/resty/v2"
)

const (
	templatePath = "internal/utils/templates/url.tpl"
)

type Agent struct {
	Client             *resty.Client
	Endpoint           string
	PollInterval       time.Duration
	UpdatePerIteration int
	MemStats           runtime.MemStats
	Metrics            metrics.Collection
}

type Config struct {
	Endpoint           string
	Timeout            time.Duration
	PollInterval       time.Duration
	UpdatePerIteration int
	Metrics            metrics.Collection
}

func NewAgent(cfg Config) *Agent {
	return &Agent{
		Client:             resty.New().SetTimeout(cfg.Timeout),
		Endpoint:           cfg.Endpoint,
		PollInterval:       cfg.PollInterval,
		UpdatePerIteration: cfg.UpdatePerIteration,
		MemStats:           runtime.MemStats{},
		Metrics:            cfg.Metrics,
	}
}

func (a *Agent) RunAgent() error {
	gauge := make(map[string]float64)
	counter := make(map[string]int64)
	iter := 0
	for {
		if iter == a.UpdatePerIteration {
			for key, val := range gauge {
				err := a.uploadMetrics(metrics.GaugeType, key, fmt.Sprintf("%.3f", val))
				if err != nil {
					return fmt.Errorf("upload failed: %v", err)
				}
			}
			for key, val := range counter {
				err := a.uploadMetrics(metrics.CounterType, key, fmt.Sprintf("%d", val))
				if err != nil {
					return fmt.Errorf("upload failed: %v", err)
				}
			}
			iter = 0
		}
		err := a.Metrics.Collect(gauge, counter, &a.MemStats)
		if err != nil {
			return err
		}

		iter++
		time.Sleep(a.PollInterval)
	}
}

func (a *Agent) uploadMetrics(metricType string, metricName string, metricValue string) error {
	url, err := utils.CreateMetricURL(templatePath, map[string]any{
		"hostname": a.Endpoint,
		"type":     metricType,
		"name":     metricName,
		"value":    metricValue,
	})
	if err != nil {
		return fmt.Errorf("cannot create url: %v", err)
	}

	resp, err := a.Client.R().Post(url)
	if err != nil {
		return fmt.Errorf("failed post: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("status code not equal 200: %v", resp.StatusCode())
	}

	return nil
}
