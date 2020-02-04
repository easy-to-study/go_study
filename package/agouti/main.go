package main

import (
	"fmt"

	"github.com/sclevine/agouti"
)

func main() {
	// Chromeを利用することを宣言
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	// 自動操作
	page.Navigate("https://qiita.com/")
	fmt.Println(page.Title())
	page.Screenshot("Screenshot01.png")

	page.FindByLink("もっと詳しく").Click()
	fmt.Println(page.Title())
	page.Screenshot("Screenshot02.png")
}
