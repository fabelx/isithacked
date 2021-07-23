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

type output struct {
	Title string
	Data  string
}

func IsItHacked(target, serviceURL string) ([]output, error) {
	var outputData []output
	c := colly.NewCollector(
		colly.AllowedDomains("isithacked.com", "www.isithacked.com"),
	)
	c.OnHTML("div.col-lg-2:has(img[alt=Xmark])~div", func(el *colly.HTMLElement) {
		outputData = append(outputData, output{
			Title: el.ChildText("h3"),
			Data:  el.ChildText("p"),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting...", r.URL.String())
	})

	err := c.Visit(fmt.Sprint(serviceURL, target, ".output"))
	if err != nil {
		return nil, err
	}

	return outputData, nil

}

func Run(cfg *config.Config) {
	if cfg.Target == "" {
		log.Fatal("No target specified.")
	}

	var re *regexp.Regexp
	if cfg.IsIp {
		re = regexp.MustCompile(config.IpRegex)
	} else {
		re = regexp.MustCompile(config.DomainRegex)
	}

	if !re.MatchString(cfg.Target) {
		log.Fatal("Invalid target format provided.")
	}

	outputData, err := IsItHacked(cfg.Target, cfg.ServiceURL)

	if err != nil {
		log.Fatalf("An error occurred while processing the request. Error: %v", err)
	}

	file, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		log.Fatalf("An error occurred while marshaling data. Error: %v", err)
	}

	err = ioutil.WriteFile(cfg.Output, file, 0644)
	if err != nil {
		log.Fatalf("An error occurred while writing to file %s. Error: %v", cfg.Output, err)
	}
}
