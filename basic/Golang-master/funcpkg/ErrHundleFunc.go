package funcpkg

import (
	"fmt"
	"os"
)

// ErrHundleFunc エラー処理のサンプルです
func ErrHundleFunc() {

	// 標準関数はほとんどが2つ目の戻り値にエラー情報を返します。
	file, err := os.Open("dummy.txt")

	if err != nil {
		fmt.Println("ファイルOPENエラー -> ", err)
	} else {
		fmt.Println("ファイルOPEN -> ", file)
	}

}
