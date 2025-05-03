package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var links []string

func getUrls(c *colly.Collector, url string) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

	})

	c.OnRequest(func(r *colly.Request) {
		myUrl := r.URL.String()
		fmt.Println(myUrl)
		links = append(links, myUrl)
	})
}
func newCollector(maxDepth int) *colly.Collector {
	collector := colly.NewCollector()
	collector.MaxDepth = maxDepth
	collector.Async = true
	return collector
}
func main() {
	c := newCollector(5)
	getUrls(c, "https://www.freecodecamp.org/news/tag/nodejs/")
	c.Visit("https://www.freecodecamp.org/news/tag/nodejs/")
	c.Wait()
	fmt.Println(links)
}
