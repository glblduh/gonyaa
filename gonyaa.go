package gonyaa

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
)

type NyaaClient struct {
	url	string
	colly	*colly.Collector
}

type NyaaResults struct {
	Category	string
	Name		string
	Size		string
	Date		string
	Seeders		int
	Leechers	int
	Downloads	int
	Magnet		string
}

func NewClient(URL string) (*NyaaClient, error) {
	c := colly.NewCollector()

	_, urlErr := url.ParseRequestURI(URL)
	if urlErr != nil {
		return nil, errors.New("Invalid URL")
	}

	return &NyaaClient{
		colly: c,
		url: URL,
	}, nil
}

func (c *NyaaClient) Search(query string, parameters string) []NyaaResults {
	var results []NyaaResults
	c.colly.OnHTML("tr", func(e *colly.HTMLElement) {
		var tempRes NyaaResults
		categoryname := e.ChildAttrs("a", "title")
		if len(categoryname) == 0 {
			return
		}
		tempRes.Category = categoryname[0]
		tempRes.Name = categoryname[len(categoryname)-1]
		e.ForEach(".text-center", func(i int, f *colly.HTMLElement) {
			switch(i) {
				case 0:
					torrentmagnet := f.ChildAttrs("a", "href")
					tempRes.Magnet = torrentmagnet[1]
				case 1:
					tempRes.Size = f.Text
				case 2:
					tempRes.Date = f.Text
				case 3:
					num, _ := strconv.Atoi(f.Text)
					tempRes.Seeders = num
				case 4:
					num, _ := strconv.Atoi(f.Text)
					tempRes.Leechers = num
				case 5:
					num, _ := strconv.Atoi(f.Text)
					tempRes.Downloads = num
			}
		})
		results = append(results, tempRes)
	})
	c.colly.Visit(c.url + "?q=" + query + parameters)
	return results
} 
