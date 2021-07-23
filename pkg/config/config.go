package config

import (
	"flag"
)

const (
	IpRegex     = `^([0-9]{1,3}\.){3}([0-9]{1,3})$`
	DomainRegex = `^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+$`
	ServiceURL  = "https://www.isithacked.com/check/"
)

var (
	Output string
	Target string
)
var IsIp bool

type Config struct {
	IsIp       bool
	Output     string
	Target     string
	ServiceURL string
}

func Init() {
	flag.StringVar(&Target, "target", "", "Target: ip (1.1.1.1) or domain (example.com)")
	flag.BoolVar(&IsIp, "ip", false, "Is target an ip?")
	flag.StringVar(&Output, "output", "output.json", "Path where output will be stored")
	flag.Parse()
}
