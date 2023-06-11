package main

import (
	"fmt"

	"github.com/XellarReps/metricscollector/internal/config"
)

func main() {
	cfg, err := config.NewConfig("config/server.yaml")
	if err != nil {
		fmt.Printf("invalid config : %v\n", err)
		return
	}

	cfg.Server.RegisterHTTP()

	err = cfg.Server.RunServer()
	if err != nil {
		panic(err)
	}
}
