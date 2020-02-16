package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sataga/go_study/qiita/excel/xlsx/sample3/config"
	"github.com/sataga/go_study/qiita/excel/xlsx/sample3/model"
	"github.com/tealeg/xlsx" // 実装した設定パッケージの読み込み
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
	fmt.Printf("r= %v , i= %v , j = %v \n", r, i, j)
	for _, cell := range row.Cells {
		fmt.Printf("cell is %v \n", cell)
		for _, targetCol := range targetColArray {
			fmt.Printf("r= %v , i= %v , j = %v \n", r, i, j)
			toInt, _ := strconv.Atoi(targetCol)
			if i == toInt {
				text := cell.String()
				fmt.Println(text)
				newsheet.Cell(r, j).Value = text
				j++
			}
		}
		i++
	}
}

func main() {
	// 設定ファイルを読み込む
	confDB, err := config.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 設定値から接続文字列を生成
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)

	// データベース接続
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	// deferで処理終了前に必ず接続をクローズする
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功！")
	}

	// // INSERTの実行
	// id, err := model.InsertUser("USER-0001", "山田 hoge郎", "pass", db)
	// if err != nil {
	//     fmt.Println(err.Error())
	// }
	// fmt.Printf("登録されたユーザのIDは【%d】です。\n", id)

	// SELECTの実行
	user, err := model.SelectUserById(1, db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("SELECTされたユーザ情報は以下の通りです。\n")
	fmt.Printf("[ID] %s\n", user.Id)
	fmt.Printf("[アカウント] %s\n", user.Account)
	fmt.Printf("[名前] %s\n", user.Name)
	fmt.Printf("[パスワード] %s\n", user.Passwd)
	fmt.Printf("[登録日] %s\n", user.Created)

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
