package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func fetch(url string) (re string, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 Edg/100.0.1185.50"
	req.Header.Set("User-Agent", userAgent)
	cookie := "_ga=GA1.2.763474059.1638023956; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8AuMt_3FvyxIgNOR82PHE4katZLabN8skZ2a6OlFwzIkTVWlq_T0UFcmKTX4F471RYlCyh_yBnqcgxTpiyrDhBLdtURm2MtAOdccWPwpJtXMv_bLxnGYQnjYUe7BcqG-1RHzwewsFxirlZ4TloYQjyY; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1649168784; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1650772194; _gid=GA1.2.1940312851.1650772195; _gat_gtag_UA_476124_1=1"
	req.Header.Add("Cookie", cookie)

	var resp *http.Response
	resp, err = client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		re = "Http get failed"
		return
	}
	if resp.StatusCode != 200 {
		re = "http status code:" + string(resp.StatusCode)
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		re = "Read Body failed"
		return
	}

	re = string(body)
	return
}

func parse(baseUrl string, html string) {
	// 删去换行符
	html = strings.Replace(html, "\n", "", -1)
	// 筛选出边栏内容块
	reSidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	// 找到边栏内容块
	sidebar := reSidebar.FindString(html)
	// 链接正则
	reLink := regexp.MustCompile(`href="(.*?)"`)
	// 找到所有的链接
	links := reLink.FindAllString(sidebar, -1)
	fmt.Println(links)

	for i, v := range links {
		// 最后三个链接是项目信息，不相关，因此不提取
		if i >= len(links)-3 {
			continue
		}
		// 把 href="" 引号中间的url 提取出来
		s := v[6 : len(v)-1]
		// 拼接完整的 url
		subUrl := baseUrl + s
		//fmt.Printf("url: %v", subUrl)
		//fmt.Println()

		body, err := fetch(subUrl)
		if err != nil {
			fmt.Printf("fetch sub page failed, err:%v\n", err)
			return
		}

		go extractArticle(body)
	}
}

// extractArticle 提取文章和标题
func extractArticle(body string) {
	// 替换空格
	body = strings.Replace(body, "\n", "", -1)
	// 获取页面内容的正则表达式
	reContent := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	// 找到页面内容
	content := reContent.FindString(body)
	// fmt.Printf("content: %v\n", content)

	// 获取标题的正则表达式
	reTitle := regexp.MustCompile(`<h1 class="article-title" itemprop="name"(.*?)</h1>`)
	// 找到标题内容
	title := reTitle.FindString(content)
	//fmt.Printf("not extracted title: %v\n", title)
	// 切片
	title = title[42 : len(title)-5]
	//fmt.Printf("title: %v\n", title)

	save(title, content)
	fmt.Println(title+".html", "保存成功！")

}

func save(title string, content string) {
	err := os.WriteFile("./html/"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "https://gorm.io/zh_CN/docs/"
	re, err := fetch(url)
	if err != nil {
		fmt.Printf("failed,reponse%s err:%v\n", re, err)
		return
	}
	// fmt.Printf("success! %s", re)
	fmt.Println("success!")
	// 对获取到的 html 做解析
	parse(url, re)
}
