package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"net/http"
)

func cks(cookie *http.Cookie) {
	name := cookie.Name
	value := cookie.Value

	fmt.Print(name + value)

}

func main() {

	url := "https://www.baidu.com"

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",                         //浏览器不提供可视化页面
			"--disable-gpu",                      //谷歌文档提到需要加上这个属性来规避bug
			"--no-sandbox",                       //不使用沙盒
			"--hide-scrollbars",                  //隐藏滚动条, 应对一些特殊页面
			"blink-settings=imagesEnabled=false", //不显示图片
			"--disable-dev-shm-usage",
			"--lang=zh-CN,zh;q=0.9",
			//"--window-size=1280,1024",
			"--disable-extensions", //不使用扩展
			"--user-agent=Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36", //自定义头部
			//"--proxy-server=http://127.0.0.1:1080", //使用proxy
			//"--user-data-dir=C:/Users/jin-zhe-hu/AppData/Local/Google/Chrome/User Data/Default", //windows os
			"--user-data-dir=/root/.config/google-chrome/", //linux os

		},
		),
		agouti.Debug,
	)

	err := driver.Start()
	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	//expire := time.Now().AddDate(0, 0, 1)
	//cookie := http.Cookie{
	//	Name:    "kims",
	//	Value:   "kory",
	//	Expires: expire,
	//}

	//cks(&cookie)
	//err = page.SetCookie(&cookie)

	err = page.Navigate(url)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	//fmt.Printf(page.Title())
	//fmt.Printf("\n")

	//html, err := page.HTML()
	//fmt.Printf (html)
	//fmt.Printf("\n")

	cookies, err := page.GetCookies()
	for _, value := range cookies {
		//fmt.Printf("  %s: %s\n", key, value)
		fmt.Printf("%s; \n", value)
	}
	//fmt.Print(cookies)

	driver.Stop()
}
