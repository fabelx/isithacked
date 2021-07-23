package main

import (
	"github.com/fabelx/isithacked/pkg/config"
	"github.com/fabelx/isithacked/pkg/isithacked"
)

func init() {
	config.Init()
}

func main() {
	cfg := &config.Config{
		Target: config.Target,
		Output: config.Output,
		IsIp:   config.IsIp,
	}
	isithacked.Run(cfg)
}
