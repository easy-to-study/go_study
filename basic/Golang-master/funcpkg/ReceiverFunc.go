package funcpkg

import "fmt"

// Tenki 型宣言
type Tenki string

// ReceiverFunc レシーバーのサンプルです
func ReceiverFunc() {
	var tk Tenki = "晴れ"
	tk.receiver()
}

func (t Tenki) receiver() {
	str := "あしたは%sです。\n"
	fmt.Printf(str, t)
}