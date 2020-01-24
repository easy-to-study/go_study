package funcpkg

import (
	"fmt"
	"time"
)

// GoroutinFunc ゴルーチンのサンプルです
// ポイント
// 関数の前にgoを付けるだけで並列処理を実装できます。
// ゴルーチン（並列処理）はとても軽いようです。
// mainの処理はゴルーチンの終了を待たない。
// ゴルーチンの結果を取得するにはチャネルを使う(チャネルは次で説明します)
func GoroutinFunc() {

	fmt.Println("GoroutinFunc 開始")

	fmt.Println("関数を実行します")
	sampleFunc()

	fmt.Println("ゴルーチンで関数を実行します")
	go sampleFunc()
	time.Sleep(1 * time.Second) // すぐ終わっちゃうのでゴルーチン用に1秒待つ

	fmt.Println("GoroutinFunc 終了")

}

func sampleFunc() {
	functions := []string{"処理1", "処理2", "処理3", "処理4"}
	for _, val := range functions {
		fmt.Println(val)
		time.Sleep(1 * time.Second)
	}
}
