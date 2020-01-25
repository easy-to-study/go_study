package main

import (
	"database/sql"
	"fmt"
	"github.com/easy-to-study/go_study/practice/go-training/conf"
	// "github.com/easy-to-study/go_study/practice/go-training/query"
	_ "github.com/go-sql-driver/mysql"
    "html/template"
    "net/http"
)
 
func main() {
 
    // 設定ファイルを読み込む
    confDB, err := conf.ReadConfDB()
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
 
    // 実践-mysql編
    // INSERTの実行
    // id, err := query.InsertUser("USER-0001", "山田 hoge郎", "pass", db)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
    // fmt.Printf("登録されたユーザのIDは【%d】です。\n", id)
 
    // SELECTの実行
    // user, err := query.SelectUserById(id, db)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
    // fmt.Printf("SELECTされたユーザ情報は以下の通りです。\n")
    // fmt.Printf("[ID] %s\n", user.Id)
    // fmt.Printf("[アカウント] %s\n", user.Account)
    // fmt.Printf("[名前] %s\n", user.Name)
    // fmt.Printf("[パスワード] %s\n", user.Passwd)
    // fmt.Printf("[登録日] %s\n", user.Created)

    // 実践-テンプレート編
    // ルートへのリクエストを"gtplHandler"関数で処理する
    http.HandleFunc("/", gtplHandler)
 
    // localhost:8080でサーバー処理開始
    http.ListenAndServe(":8080", nil)
}

func gtplHandler(w http.ResponseWriter, r *http.Request) {
 
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("templates/sample.gtpl"))
 
    // テンプレートに出力する値をマップにセット
    values := map[string]string{
        "account": "user-0001",
        "name":    "山田太郎",
        "passwd":  "sample-pass",
    }
 
    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "sample.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
