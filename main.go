package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	r := gin.Default()

	v := url.Values{}

	key_url := "https://www.creditchina.gov.cn/api/public_search/getCreditCodeFacades?"
	end_url := "&filterManageDept=0&filterOrgan=0&filterDivisionCode=0&page=1&pageSize=1"

	client := &http.Client{}

	r.GET("/key/:name", func(c *gin.Context) {

		name := c.Param("name")

		v.Add("keyword", name)
		keys := v.Encode()

		req, err := http.NewRequest("GET", key_url+keys+end_url, nil)
		if err != nil {
			log.Fatal(err)
		}

		//req.Header.Set("Connection", "keep-alive")
		//req.Header.Set("Pragma", "no-cache")
		//req.Header.Set("Cache-Control", "no-cache")
		//req.Header.Set("Upgrade-Insecure-Requests", "1")
		//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
		//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		cn_json, _ := simplejson.NewJson([]byte(bodyText))
		//cn_results, _ := cn_json.Get("results").Array()
		//cn_page, _ := cn_json.Get("page").Int()
		//cn_message, _ := cn_json.Get("message").String()

		//for _, v := range cn_results {
		//	fmt.Println(v)
		//}

		//c.JSON(200, gin.H{
		//	"JSON_results": cn_json,
		//})

		c.JSONP(http.StatusOK, cn_json)

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
