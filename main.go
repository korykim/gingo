package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	r := gin.Default()

	v := url.Values{}
	v.Add("keyword","洋光东森(北京)商贸有限公司")
	keys :=v.Encode()

	key_url :="https://www.creditchina.gov.cn/api/public_search/getCreditCodeFacades?"
	end_url :="&filterManageDept=0&filterOrgan=0&filterDivisionCode=0&page=1&pageSize=1"

	client := &http.Client{}

	r.GET("/ping", func(c *gin.Context) {
		req, err := http.NewRequest("GET", key_url+keys+end_url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", bodyText)

		c.JSON(200, gin.H{
			"message": "ok",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
