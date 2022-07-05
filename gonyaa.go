package gonyaa

import (
	"fmt"
	"github.com/gocolly/colly"
)

type NyaaClient struct {
	Colly	*colly.Collector
}

type NyaaResults struct {
	Catergory	string
	Name		string
	Size		string
	Date		string
	Seeders		int
	Leechers	int
	Downloads	int
}

func NewClient(url string) *NyaaClient {
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)
	return &NyaaClient{
		Colly: c,
	}
}

func (c *NyaaClient) Search(query string, parameters string) NyaaResults {
	c.Colly.OnHTML("tr", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	return NyaaResults{}
} 
