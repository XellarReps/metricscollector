package main

import (
	"fmt"
	"github.com/XellarReps/metricscollector/internal/config"
)

const (
	configPath = "config/metricscollector.yaml"
)

func main() {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		fmt.Printf("invalid config : %v\n", err)
		return
	}

	err = cfg.Agent.RunAgent()
	if err != nil {
		panic(err)
	}
}
