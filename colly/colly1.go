package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("hackerspaces.org"),
	)

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", element.Text, link)
		c.Visit(element.Request.AbsoluteURL(link))
	})

	//伪装成真实浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.k2c.link")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Pragma", "no-cache")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")

		fmt.Println("Visiting", r.URL.String())

	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit("https://hackerspaces.org/")

}
