package main

import (
	"github.com/fabelx/isithacked/pkg/app"
	"github.com/fabelx/isithacked/pkg/config"
)

func init() {
	config.Init()
}

func main() {
	cfg := &config.Config{
		Target:     config.Target,
		Output:     config.Output,
		IsIp:       config.IsIp,
		ServiceURL: config.ServiceURL,
	}
	app.Run(cfg)
}
