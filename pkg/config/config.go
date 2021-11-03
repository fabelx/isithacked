package config

import (
	"flag"
)

const (
	DomainRegex = `^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+$`
	ServiceURL  = "https://www.isithacked.com/check/"
)

var (
	Output string
	Domain string
)

type Config struct {
	Output string
	Domain string
}

func Init() {
	flag.StringVar(&Domain, "target", "", "Domain (example.com)")
	flag.StringVar(&Output, "output", "output.json", "Path where output will be stored")
	flag.Parse()
}
