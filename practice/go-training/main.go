package main

import (
	"net/http"
	// 実装したHTTPリクエストハンドラパッケージの読み込み
	"github.com/easy-to-study/go_study/practice/go-training/req_handler"
)

func main() {

	// "user-list"へのリクエストを関数で処理する
	http.HandleFunc("/user-list", req_handler.HandlerUserList)

	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/user-form", req_handler.HandlerUserForm)

	// "user-confirm"へのリクエストを関数で処理する
	http.HandleFunc("/user-confirm", req_handler.HandlerUserConfirm)

	// "user-registered"へのリクエストを関数で処理する
	http.HandleFunc("/user-registered", req_handler.HandlerUserRegistered)

	// "user-edit"へのリクエストを関数で処理する
	http.HandleFunc("/user-edit", req_handler.HandlerUserEdit)

	// "user-update"へのリクエストを関数で処理する
	http.HandleFunc("/user-update", req_handler.HandlerUserUpdate)

	// css・js・イメージファイル等の静的ファイル格納パス
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset/"))))

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}
