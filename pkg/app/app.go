package app

import (
	"fmt"
	"github.com/fabelx/isithacked/pkg/config"
	"io"
	"log"
	"net/http"
	"regexp"
)

var re *regexp.Regexp

func parseResponse(body []byte) {

}

func Run(cfg *config.Config) {
	if cfg.Target == "" {
		log.Fatal("No target specified.")
	}

	if cfg.IsIp {
		re = regexp.MustCompile(config.IpRegex)
	} else {
		re = regexp.MustCompile(config.DomainRegex)
	}

	if !re.MatchString(cfg.Target) {
		log.Fatal("Invalid target format provided.")
	}

	resp, err := http.Get(fmt.Sprint(config.ServiceURL, cfg.Target))
	if err != nil {
		log.Fatalf("An error occurred while processing the request. Error: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	parseResponse(body)
}
