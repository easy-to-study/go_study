package funcpkg

import (
	"fmt"
	"strconv"
)

// ForFunc for文のサンプルです
func ForFunc() {

	for i := 0; i < 3; i++ {
		fmt.Println("i -> " + strconv.Itoa(i))
	}

	// whileのかわり
	j := 3
	for {
		if j > 5 {
			break
		}

		fmt.Println("j -> " + strconv.Itoa(j))
		j++
	}

	// foreach
	// "_"にindex値、"k"に配列の値が入る
	kList := [...]string{"6", "7", "8"}
	for _, k := range kList {
		fmt.Println("k -> " + k)
	}

}