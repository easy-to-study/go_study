package funcpkg

import "fmt"

// AboutFunc 関数のサンプルです
func AboutFunc() {
	args("A", "B", "C", "D")                            // 可変長引数
	fmt.Printf("simpleReturn -> %d\n", simpleReturn(1)) // 戻り値を返す

	ret1, ret2 := multiReturn(1, 2)
	fmt.Printf("multiReturn  -> %d, %d\n", ret1, ret2) // 複数の戻り値を返す
}

// 可変長引数
func args(vals ...string) {
	fmt.Print("args -> ")
	for _, val := range vals {
		fmt.Print(val + " ")
	}
	fmt.Println()
}

// 戻り値を返す
func simpleReturn(num int) int {
	return num
}

// 複数の戻り値を返す
func multiReturn(num1 int, num2 int) (int, int) {
	return num1 * 2, num2 * 2
}