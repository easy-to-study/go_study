package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile("Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				v := cell.String()
				fmt.Println(v)
			}
		}
	}
}

// OpenFile("test.xlsx")でExcelファイルを読み込む。
// for _, sheet := range file.Sheets でファイルの中のシートを順に見て、
// for _, row := range sheet.Rows で、上の行から順に見て、
// for _, cell := range row.Cells で列をAからZへと左方向へ見ていく。
// 列の中にデータがあれば、v := cell.String()でvにセルの値を文字列のデータ
