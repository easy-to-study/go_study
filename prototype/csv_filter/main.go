package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type List struct {
	FileName  string
	StartRow  int
	SheetName string
	SearchCol int
	TargetCol string
	ItemIndex int
	OutPut    string
}

type ResultSet struct {
	List []List
}

var conf ResultSet

func main() {
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
		//行を新ファイルに追加
		if cnt >= startRow {
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
	for _, cell := range row.Cells {
		fmt.Printf("r=%v / cell=%v \n", r, cell)
		for _, targetCol := range targetColArray {
			toInt, _ := strconv.Atoi(targetCol)
			fmt.Printf("targetCol=%v / i=%v , j=%v \n", targetCol, i, j)
			text := cell.String()
			if i == toInt {
				switch i {
				case 0:
					if text != "stg" {
						newsheet.Cell(r, j).Value = text
					}
					j++
				case 2:
					if text == "kaihatsu" || text == "verify" {
					} else {
						newsheet.Cell(r, j).Value = text
					}
					j++
				case 4:
					if text == "TRUE" {
						newsheet.Cell(r, j).Value = "有"
					} else if text == "FALSE" {
						newsheet.Cell(r, j).Value = "無"
					} else {
						newsheet.Cell(r, j).Value = text
					}
					j++
				default:
					fmt.Println("column=other")
				}

			}
		}
		i++
	}
}
