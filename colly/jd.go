package main

import "github.com/gocolly/colly"

const goodsUrl = "https://item.jd.com/100026667872.html"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.jd.com"),
	)

	//c.OnHTML()

	c.Visit(goodsUrl)
}
