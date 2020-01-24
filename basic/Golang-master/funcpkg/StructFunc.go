package funcpkg

import "fmt"

// Vector 構造体
type Vector struct {
	structA string
	structB string
}

// StructFunc 構造体のサンプルです
func StructFunc() {
	s := Vector{structA: "AAA", structB: "BBB"}
	fmt.Println("Vector -> ", s)
}
