package main

import (
	"github.com/fabelx/isithacked/pkg/app"
	"github.com/fabelx/isithacked/pkg/config"
)

var cfg *config.Config

func init() {
	cfg = config.Init()
}

func main() {
	app.Run(cfg)
}
