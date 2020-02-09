package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	// chromeの起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	// google.comへ
	if err := page.Navigate("https://www.google.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	// 入力
	q := page.FindByName("q")
	q.Fill("乃木坂46")

	// 検索
	if err := q.Submit(); err != nil {
		log.Fatalf("Failed to Search:%v", err)
	}

	// スクショ
	if err := page.Screenshot("screenshot.png"); err != nil {
		log.Fatalf("Failed to Screenshot:%v", err)
	}
}
