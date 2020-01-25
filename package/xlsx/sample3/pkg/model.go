package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() {

	// データベース接続
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/sampledb?charset=utf8")
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
}
