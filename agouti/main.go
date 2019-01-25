package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
)

func main() {

	fmt.Fprintf(os.Stderr, "*** Start***\n")

	url := "https://www.facebook.com" //http://localhost:5000/head

	//noinspection ALL
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",                         //浏览器不提供可视化页面
			"--disable-gpu",                      //谷歌文档提到需要加上这个属性来规避bug
			"--no-sandbox",                       //不使用沙盒
			"--hide-scrollbars",                  //隐藏滚动条, 应对一些特殊页面
			"blink-settings=imagesEnabled=false", //不显示图片
			"--disable-dev-shm-usage",
			//"--window-size=1280,1024",
			"--disable-extensions", //不使用扩展
			"--user-agent=Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
			"--proxy-server=http://127.0.0.1:1080",
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

	err = page.Navigate(url)

	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	//page.Screenshot("xin.jpg")
	//code,err  := page.Find("#basic-wrap").Text()
	//Ecode,err := page.Find("#basic-wrap").Elements()

	//code,err  := page.Find("#basic-wrap > div > div.zx-detail-basic-info > div:nth-child(2) > p:nth-child(1) > span").Text()

	//if err != nil {
	//	log.Printf("Failed to navigate: %v", err)
	//}

	//fmt.Print(reflect.TypeOf(Ecode).String())

	//fmt.Printf(code)
	//fmt.Printf("\n")

	//for index,value := range Ecode{
	//	fmt.Println(index, *value)
	//}

	html, err := page.HTML()
	fmt.Printf(html)
	fmt.Printf("\n")

	//time.Sleep(100 * time.Millisecond)
	//
	//out_filename := "tmp001.html"
	//ioutil.WriteFile (out_filename,[]byte(html),0666)

	fmt.Fprintf(os.Stderr, "*** End ***\n")

}
