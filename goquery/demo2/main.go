package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)

	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		DetailUrl := url + link
		fmt.Printf("detail_url: %v\n", DetailUrl)
		d, _ = goquery.NewDocument(DetailUrl)

		title := d.Find(".article-title").Text()
		content, _ := d.Find(".article").Html()

		fmt.Printf("title: %v\n", title)
		fmt.Printf("content: %v \n", content)
	})

}
