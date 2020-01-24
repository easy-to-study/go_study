package funcpkg

import (
	"fmt"
)

// RecoveryFunc recoveryのサンプルです
// ポイント
// panicが発生した時だけ処理したい場合、recoveryを使います。
func RecoveryFunc() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("パニックが発生してるのでここでリカバリー")
		}
	}()

	fmt.Println("処理 start")

	panic("パニック発生！！")

	fmt.Println("処理 end")
}
