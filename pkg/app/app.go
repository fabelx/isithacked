package app

import (
	"encoding/json"
	"fmt"
	"github.com/fabelx/isithacked/pkg/config"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"regexp"
)

type Output struct {
	Title string
	Data  string
}

var re *regexp.Regexp
var outputData []Output

func IsItHacked(target string, output string) {
	url := fmt.Sprint(config.ServiceURL, target, ".output")
	c := colly.NewCollector(
		colly.AllowedDomains("isithacked.com", "www.isithacked.com"),
	)
	c.OnHTML("div.col-lg-2:has(img[alt=Xmark])~div", func(el *colly.HTMLElement) {
		outputData = append(outputData, Output{
			Title: el.ChildText("h3"),
			Data:  el.ChildText("p"),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting...", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatalf("An error occurred while processing the request. Error: %v", err)
	}

	file, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		log.Fatalf("An error occurred while marshaling data. Error: %v", err)
	}

	err = ioutil.WriteFile(output, file, 0644)
	if err != nil {
		log.Fatalf("An error occurred while writing to file %s. Error: %v", output, err)
	}
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

	IsItHacked(cfg.Target, cfg.Output)
}
