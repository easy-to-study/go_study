package main

import (
	"encoding/json"
	"fmt"
)

// type Musician struct {
// 	Name       string `json:"name"`
// 	Instrument string `json:"instrument"`
// }

// type XJapan struct {
// 	Members []Musician `json:"members"`
// 	Songs   []string   `json:"songs"`
// }

// 構造体を入れ子にできる
// 個別の構造体を定義した場合と同じ意味！
type XJapan struct {
	Members []struct {
		Name       string `json:"name"`
		Instrument string `json:"instrument"`
	} `json:"members"`
	Songs []string `json:"songs"`
}

func main() {
	var xjapanJson string = `
	{
    "members": [
        { "name": "Toshl",   "instrument": "vocal"  },
        { "name": "PATA",    "instrument": "guitar" },
        { "name": "HEATH",   "instrument": "bass"   },
        { "name": "SUGIZO",  "instrument": "guitar" },
        { "name": "YOSHIKI", "instrument": "drums"  }
    ],
    "songs": ["紅", "Silent Jealousy"]
	}
	`
	// Unmarshal結果の格納先である構造体のポインターを取得
	xJapan := new(XJapan)

	// JSON文字列をバイト列にキャスト
	jsonBytes := []byte(xjapanJson)

	// xJapanにバイト列を格納する
	if err := json.Unmarshal(jsonBytes, xJapan); err != nil {
		fmt.Println(err)
		return
	}
	for _, members := range xJapan.Members {
		fmt.Printf("NAME: %-7s INSTRUMENT: %s\n", members.Name, members.Instrument)
	}
	/*
	   出力結果
	   NAME: Toshl   INSTRUMENT: vocal
	   NAME: PATA    INSTRUMENT: guitar
	   NAME: HEATH   INSTRUMENT: bass
	   NAME: SUGIZO  INSTRUMENT: guitar
	   NAME: YOSHIKI INSTRUMENT: drums
	*/

	for _, song := range xJapan.Songs {
		fmt.Println(song)
	}
	/*
	   出力結果
	   紅
	   Silent Jealousy
	*/
}
