package gonyaa

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

type NyaaClient struct {
	url	string
	colly	*colly.Collector
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
		colly: c,
	}
}

func (c *NyaaClient) Search(query string, parameters string) NyaaResults {
	c.colly.OnHTML("tr", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)	
	})
	c.colly.Visit("https://" + c.url + "?q=" + url.QueryEscape(query) + parameters)
	return NyaaResults{}
} 
