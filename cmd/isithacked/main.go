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
		Domain: config.Domain,
		Output: config.Output,
	}
	isithacked.Run(cfg)
}
