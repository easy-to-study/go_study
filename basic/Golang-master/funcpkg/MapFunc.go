package funcpkg

import "fmt"

// MapFunc マップのサンプルです
// ポイント
// マップの格納順は補償されない
// 存在しないキーを指定してもエラーにならずnilが返る
func MapFunc() {

	// マップ作成
	enviroment := map[string]string{
		"OS":   "Windows",
		"Lang": "Golang",
		"Web":  "nginx",
	}

	// 個別に値セット
	enviroment["DB"] = "mongoDB"

	// 格納順は補償されない
	fmt.Println("マップ全て     -> ", enviroment)
	fmt.Println("キー指定       -> ", enviroment["Lang"])
	fmt.Println("存在しないキー  -> ", enviroment["Framework"])

	// キーの削除
	delete(enviroment, "Web")

	// キーの存在確認
	key, exist := enviroment["Web"]
	if exist {
		fmt.Println(key, "は存在する")
	} else {
		// ここにくる場合は、keyはnil
		fmt.Println("Webは存在しない")
	}
}
