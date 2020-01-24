package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/urfave/cli"
)

func main() {
	// アプリケーションを作成
	app := cli.NewApp()
	app.Name = "my_cli"
	app.Usage = "my commnd line application"
	// (単発であればこれでいい)アプリケーションの実行内容
	// app.Action = func(c *cli.Context) error {
	// 	fmt.Println("boom! I say!")
	// 	return nil
	// }

	app.Commands = []cli.Command{
		{
			// qrcode
			Name:    "qrcode",
			Aliases: []string{"q"},
			Usage:   "make qrcode",
			Action: func(c *cli.Context) error {
				// cには、qrcodeの後の文字列が取得される
				fmt.Println("qrcode complate: ", c.Args().First())
				// Create the barcode
				qrCode, _ := qr.Encode(c.Args().First(), qr.M, qr.Auto)

				// Scale the barcode to 200x200 pixels
				qrCode, _ = barcode.Scale(qrCode, 200, 200)

				// create the output file
				file, _ := os.Create("qrcode.png")
				defer file.Close()

				// encode the barcode as png
				png.Encode(file, qrCode)
				return nil
			},
		},
		{
			// webのatabのcrawl作成
			Name:    "crawler",
			Aliases: []string{"c"},
			Usage:   "crawler atags",
			Action: func(c *cli.Context) error {
				// Request the HTML page.
				res, err := http.Get(c.Args().First())
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
					title := s.Find("i").Text()
					fmt.Printf("Review %d: %s - %s\n", i, link, title)
				})
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
