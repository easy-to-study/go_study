package funcpkg

import "fmt"

// IfFunc if文のサンプルです
func IfFunc() {
	num := 5 // 変数の宣言と値の設定をまとめて

	if num < 10 {
		fmt.Println("num is small!!")
	} else if num > 10 {
		fmt.Println("num is big!!")
	} else {
		fmt.Println("punic!!")
	}

}