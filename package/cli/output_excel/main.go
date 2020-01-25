package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/tealeg/xlsx"

)

func main() {

	var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
    }

	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a")
		link, _ := band.Last().Attr("href")
		// コンソールに出力するだけ
		// fmt.Printf("Review %d: %s - %s\n", i, link)

		// Excelデータの中に新しい行を追加していく
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = link
		err = file.Save("MyXLSXFile.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	})
}
