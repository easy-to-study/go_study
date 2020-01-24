package funcpkg

import (
	"fmt"
)

// PanicFunc panicのサンプルです
// ポイント
// panicが発生すると以降のコードは実行されません。
// 但し、deferはpanicが発生しても実行されます。
// またdeferはそれぞれの関数の最後に実行されます。
func PanicFunc() {
	// deferはpanicが発生しても呼ばれます、ファイルやDBのクローズ処理を書くと良いでしょう
	defer func() {
		fmt.Println("後処理２")
	}()

	fmt.Println("PanicFunc 開始")

	doPanicFunc()

	fmt.Println("PanicFunc 終了") // ここは実行されない
}

func doPanicFunc() {
	defer func() {
		fmt.Println("後処理１")
	}()
	// 何かいろいろ処理してパニック発生する
	panic("パニック発生")
}
