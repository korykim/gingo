package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
)

func main() {

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",                         //浏览器不提供可视化页面
			"--disable-gpu",                      //谷歌文档提到需要加上这个属性来规避bug
			"--no-sandbox",                       //不使用沙盒
			"--hide-scrollbars",                  //隐藏滚动条, 应对一些特殊页面
			"blink-settings=imagesEnabled=false", //不显示图片
			"--disable-dev-shm-usage",
			"--disable-translate", //不使用翻译
			//"--window-size=1280,1024",
			"--disable-extensions", //不使用扩展
			"--user-agent=Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
			"--lang=zh-CN,zh;q=0.9",
			//"--user-data-dir=C:/Users/jin-zhe-hu/AppData/Local/Google/Chrome/User Data/Default", //windows os
			"--user-data-dir=/root/.config/google-chrome/", //linux os

		},
		),
		agouti.Debug,
		agouti.Browser("chrome"),
	)

	proxys := agouti.ProxyConfig{
		ProxyType:          "pac", //direct,manual,pac,autodetect,system
		ProxyAutoconfigURL: "http://127.0.0.1:1080/pac?t=20190121081039383&secret=utMghubXpHJ8AbBqzUVuEBQ073t/cu0jvg7Ha540fUg=",
		//HTTPProxy:          "http://127.0.0.1:1080",
		//SSLProxy:           "http://127.0.0.1:1080",
		//SOCKSProxy:         "http://127.0.0.1:1080",
	}

	capabilities := agouti.NewCapabilities().Browser("chrome").Proxy(proxys).Without("javascriptEnabled")
	err := driver.Start()

	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}

	page, err := driver.NewPage(agouti.Desired(capabilities))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate("https://www.1688.com")
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	cookies, err := page.GetCookies()
	fmt.Print(cookies)

	//html, err := page.HTML()
	//fmt.Printf (html)
	//fmt.Printf("\n")
	fmt.Printf(page.Title())
	fmt.Printf("\n")

	driver.Stop()

}
