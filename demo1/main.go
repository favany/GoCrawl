package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	url := "https://movie.douban.com/top250"
	re, err := fetch(url)
	if err != nil {
		fmt.Printf("failed,reponse%s err:%v\n", re, err)
		return
	}
	fmt.Printf("success! %s", re)
}
