package funcpkg

import "fmt"

// PointerFunc ポインタのサンプルです
func PointerFunc() {
	intA := 10
	intB := 10

	calc(intA, &intB) // "&"を付ける

	fmt.Println("値渡し   -> ", intA)
	fmt.Println("ポインタ -> ", intB)
}

// "calc"のように先頭が小文字だと他ファイルからは呼べない
func calc(intA int, intB *int) {
	intA = intA + 1   // 値渡し
	*intB = *intB + 1 // ポインタ（参照渡し）
}