package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

func main() {
	var user = os.Getenv("USER")
	var url = "file:///Users/" + user + "/go/src/github.com/easy-to-study/go_study/practice/scraping/sample.html" // urlを指定
	driver := agouti.ChromeDriver()                                                                               // ドライバの起動

	err := driver.Start()
	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	// クロームを起動。page型の返り値（セッション）を返す
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate(url) // 指定したurlにアクセスする
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	curContentsDom, err := page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	readerCurContents := strings.NewReader(curContentsDom)
	contentsDom, _ := goquery.NewDocumentFromReader(readerCurContents) // 現状ブラウザで開いているページのDOMを取得
	listDom := contentsDom.Find("#target").Children()                  // selector 　部分にセレクトボックスのセレクタを入れる。セレクトボックス内の子要素を全部取得
	listLen := listDom.Length()                                        // セレクトボックスの子要素の個数を取得

	for i := 1; i <= listLen; i++ {
		iStr := strconv.Itoa(i)
		page.Find(" #target > option:nth-child(" + iStr + ")").Click()          // セレクタの属性（ここではoption）のバリューで繰り返しクリックする
		tmp, _ := page.Find(" #target > option:nth-child(" + iStr + ")").Text() //
		fmt.Println(tmp)

		time.Sleep(2 * time.Second) //適宜ブラウザが反応するための間を取る

		curContentsDom, err := page.HTML()
		if err != nil {
			log.Printf("Failed to get html: %v", err)
		}

		readerCurContents := strings.NewReader(curContentsDom)
		contentsDom, _ := goquery.NewDocumentFromReader(readerCurContents) // ページ内容が変化しているので再取得

		contentsDom.Find("#target > option ").Each(func(_ int, s *goquery.Selection) { // 繰り返し取得したい部分までのセレクタを入れる
			data := s.Text() // テキスト要素を取得したい場合
			// link, _ := s.Find("option").Attr("href") // アンカーリンク要素を取得したい場合
			// alt, _ := s.Find("option").Attr("alt")   // alt要素を取得したい場合
			fmt.Println(data)
			time.Sleep(time.Second * 1) //サーバに負荷を与えないように。
		})
	}

	// 表示確認するために待たせる
	time.Sleep(time.Second * 5)

}
