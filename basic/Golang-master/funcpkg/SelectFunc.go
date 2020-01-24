package funcpkg

import (
	"fmt"
	"math/rand"
)

// SelectFunc ゴルーチンのサンプルです
// ポイント
// 複数のチャネルを待機する場合は、select文を使う
// defaultはどのチャネルも受診しなかった時に呼ばれる
// defaultが無い場合は、受信するまで待機する
func SelectFunc()  {
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch0 := make(chan int)

	fmt.Println("開始")

	go func(c1 chan<- int, c2 chan<- int, c0 chan<- int) {

		for {
			// 0〜9のランダムな数値
			randInt := rand.Intn(10)

			if randInt == 0 {
				ch0 <- randInt
				break
			} else if randInt%2 == 0 {
				ch1 <- randInt
			} else {
				ch2 <- randInt
			}
		}

	}(ch1, ch2, ch0)

	for {
		select {
		case ch1Int := <-ch1:
			fmt.Println("ch1受信 : ", ch1Int)
		case ch2Int := <-ch2:
			fmt.Println("ch2受信 : ", ch2Int)
		case ch0Int := <-ch0:
			fmt.Println("終了    : ", ch0Int)
			return
		}
	}

}