package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
	"github.com/sataga/go_study/qiita/excel/xlsx/sample3/pkg"
)

type List struct {
	FileName   string
	StartRow   int
	SheetName  string
	SearchCol  int
	TargetCol  string
	SearchText string
	ItemIndex  int
	OutPut     string
}

type ResultSet struct {
	List []List
}

var conf ResultSet

func getConf() {
	data, _ := ioutil.ReadFile("conf.xml")
	err := xml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
}

func outputData(row *xlsx.Row, newsheet *xlsx.Sheet, r int) {
	targetColArray := strings.Split(conf.List[0].TargetCol, ",")
	fmt.Println(targetColArray)
	i := 0
	j := 0
	fmt.Println(r)
	for _, cell := range row.Cells {
		fmt.Println(i)

		for _, targetCol := range targetColArray {
			toInt, _ := strconv.Atoi(targetCol)
			if i == toInt {
				text := cell.String()
				newsheet.Cell(r, j).Value = text
				j++
			}
		}
		i++
	}
}

func main() {
	model.DbConnect()
}

func test() {
	//パラメーターをxmlファイルから取得する
	getConf()

	fmt.Println("ファイル読み込み中...")

	//Excelファイルを開く
	excel, err1 := xlsx.OpenFile(conf.List[0].FileName)
	if err1 != nil {
		fmt.Printf(err1.Error())
	}

	fmt.Println("指定シート検索中...")

	//指定シートを検索して変数sheetへ
	var sheet *xlsx.Sheet
	for _, s := range excel.Sheets {
		if s.Name == conf.List[0].SheetName {
			sheet = s
			break
		}
	}

	//新しいExcelファイルを作成
	newfile := xlsx.NewFile()
	newsheet, err := newfile.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("検索中...")

	//検索&ファイル書き込み
	//パラメーター
	index := conf.List[0].SearchCol - 1
	text := conf.List[0].SearchText
	itemIndex := conf.List[0].ItemIndex - 1
	startRow := conf.List[0].StartRow - 1
	//検索
	r := 1
	cnt := 0
	for _, row := range sheet.Rows {
		//項目行を最初に追加
		if cnt == itemIndex {
			outputData(row, newsheet, 0)
		}
		//検索ヒットした行を新ファイルに追加
		if cnt >= startRow && row.Cells[index].Value == text {
			outputData(row, newsheet, r)
			r++
		}
		cnt++
	}

	//ファイルに保存（ファイル名はパラメーターから取得）
	err = newfile.Save(conf.List[0].OutPut)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("完了！")
}
