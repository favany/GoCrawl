package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getDoc1() {
	url := "https://gorm.io/zh_CN/docs"
	d, _ := goquery.NewDocument(url)
	d.Find(".sidebar-link").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		fmt.Println("text", s)
		href, _ := selection.Attr("href")
		fmt.Printf("href: %v\n", href)
	})
}

func getDoc2() {
	client := &http.Client{}
	url := "https://gorm.io/zh_CN/docs"
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
		return
	}
	dom.Find(".sidebar-link").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		fmt.Println("text", s)
		href, _ := selection.Attr("href")
		fmt.Printf("href: %v\n", href)
	})
}

func getDoc3() {
	html := `<body>
				<div id="div1" class="c1">DIV1</div>
				<div>DIV2</div>
				<div>DIV3</div>
			</body>
	`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Printf("get dom failed, err:%v\n", err)
		return
	}

	// 标签选择器 <div></div>
	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		fmt.Println("标签选择器查找", s)
	})

	// 类选择器 class="c1"
	dom.Find(".c1").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		fmt.Println("类选择器 查找", s)
	})

	// id 选择器 id="div1"
	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		s := selection.Text()
		fmt.Println("id选择器 查找", s)
	})

}

func main() {
	//getDoc1()
	//getDoc2()
	getDoc3()
}
