package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// // ネストする場合
// type Person struct {
// 	Id       int      `json:"id"`
// 	Name     string   `json:"name"`
// 	Birthday string   `json:"birthday"`
// 	Vivid    struct { // <- 構造体の中にネストさせて構造体を定義
// 		Color  string `json:color`
// 		Weapon string `json:weapon`
// 	} `json:"vivid_info"`
// }

// // 別の構造体に分ける場合
type VividInfo struct {
	Color  string `json:color`
	Weapon string `json:weapon`
}

// type Person struct {
// 	Id       int       `json:"id"`
// 	Name     string    `json:"name"`
// 	Birthday string    `json:"birthday"`
// 	Vivid    VividInfo `json:"vivid_info"`
// }

// Birthdayをstringのポインタで定義
type Person struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Birthday *string    `json:"birthday"`   // <- ポインタ定義
	Vivid    *VividInfo `json:"vivid_info"` // <- ポインタ定義
}

func main() {
	bytes, err := ioutil.ReadFile("vro.json")
	if err != nil {
		log.Fatal(err)
	}
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}
	// 表示
	for _, p := range persons {
		fmt.Println(p.GetInfo())
	}
}

func (p *Person) GetInfo() string {
	if p.Birthday != nil {
		return fmt.Sprintf("%d : %s [%s]", p.Id, p.Name, *p.Birthday)
	}
	return fmt.Sprintf("%d : %s", p.Id, p.Name)
}
