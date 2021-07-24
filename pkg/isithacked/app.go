package isithacked

import (
	"encoding/json"
	"fmt"
	"github.com/fabelx/isithacked/pkg/config"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Output struct {
	Title string `json:"title"`
	Data  string `json:"data"`
}

func IsItHacked(target string) ([]*Output, error) {
	var outputData []*Output
	c := colly.NewCollector(
		colly.AllowedDomains("isithacked.com", "www.isithacked.com"),
	)
	c.OnHTML("div.col-lg-2:has(img[alt=Xmark])~div", func(el *colly.HTMLElement) {
		outputData = append(outputData, &Output{
			Title: el.ChildText("h3"),
			Data:  el.ChildText("p"),
		})
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting...", r.URL.String())
	})

	err := c.Visit(fmt.Sprint(config.ServiceURL, target, ".output"))
	if err != nil {
		return nil, err
	}

	return outputData, nil

}

func selectRegex(f bool) string {
	if f {
		return config.IpRegex
	}
	return config.DomainRegex
}

func Run(cfg *config.Config) {
	if cfg.Target == "" {
		log.Fatal("No target specified.")
	}

	re := regexp.MustCompile(selectRegex(cfg.IsIp))
	if !re.MatchString(cfg.Target) {
		log.Fatal("Invalid target format provided.")
	}

	outputData, err := IsItHacked(cfg.Target)
	if err != nil {
		log.Fatalf("An error occurred while processing the request. Error: %v", err)
	}

	if outputData == nil {
		log.Printf("No issues found for target: %s, Congrats!", cfg.Target)
		os.Exit(0)
	}

	var s = "s"
	if len(outputData) == 1 {
		s = ""
	}

	log.Printf("Found %v issue%s!", len(outputData), s)
	file, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		log.Fatalf("An error occurred while marshaling data. Error: %v", err)
	}

	err = ioutil.WriteFile(cfg.Output, file, 0644)
	if err != nil {
		log.Fatalf("An error occurred while writing to file %s. Error: %v", cfg.Output, err)
	}
}
