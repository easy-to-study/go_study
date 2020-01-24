package funcpkg

import "fmt"

// チャネルの宣言の仕方
// chan 型			送受信可能なチャネル
// chan <- 型		送信専用チャネル
// <- 型			受信専用チャネル

// ChanelFunc チャネルのサンプルです
// 送信専用、受信専用は主に関数の引数として受け取る際に使用する
// 使わなくなったチャネルは最後にcloseする
// 受信専用チャネルはクローズできない
// クローズ済みのチャネルにデータ送信するとランタイムパニックが発生

func ChanelFunc()  {
	msgList := [...]string{"受信", "受け取った", "Receive"}

	msgChanel := make(chan string)

	// チャネルに送信
	go func (cha chan<- string )  {
		for index := 0; index < len(msgList); index++ {
			cha <- msgList[index]
		}	
		close(cha)
	}(msgChanel)

	for index2 := 0; index2 < len(msgList); index2++ {
		getMsg := <-msgChanel
		fmt.Println(getMsg)	
	}
}