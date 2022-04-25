package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// goquery selector class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url", r.URL)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("请求前调用 onRequest")
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("发生错误调用 onError")
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("获得响应后调用 onResponse")
	})

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		fmt.Println("onResponse 收到 HTML 内容后调用: onHTML")
	})

	c.OnXML("//h1", func(element *colly.XMLElement) {
		fmt.Println("")
	})

	c.Visit("https://gorm.io/zh_CN/docs")
}
