package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// goquery selector class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))

		href := e.Attr("href")

		//fmt.Println("url", s, s2)

		if href != "index.html" {
			c.Visit(e.Request.AbsoluteURL(href))
		}
	})

	c.OnHTML(".article-title", func(element *colly.HTMLElement) {
		s := element.Text
		fmt.Println(s)
	})

	c.OnHTML(".article", func(h *colly.HTMLElement) {
		content, _ := h.DOM.Html()
		fmt.Println("Visiting", content)
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("url", r.URL)
	})

	c.Visit("https://gorm.io/zh_CN/docs")
}
