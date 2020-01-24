package funcpkg

import (
	"fmt"
	"time"
	"sync"
)

// SyncFunc 同期処理のサンプルです
// ポイント
// WaitGroupをインクリメントする
// Doneでデクリメントする
// WaitはWaitGroupが0になるまで待つ
// つまり並列でやりたい処理の数分インクリメントしてそれぞれの処理が終わればデクリメントして
// WaitGroupが0になれば全ての並列処理が完了
func SyncFunc()  {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)	

		go func(cnt int) {
			fmt.Println("start", cnt)
			time.Sleep(1 * time.Millisecond)
			fmt.Println("end", cnt)

			wg.Done()
		}(i)
	}

	wg.Wait()
}