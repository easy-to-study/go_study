package funcpkg

import "fmt"

// InterfaceFunc Interfaceのサンプルです
// ダックタイピング
// 9～11行目のインターフェースでEat()を実装してるのでCarlingGirlもFoodFighterもHumanとして扱える
// 正直Java経験者の私にはピンとこない仕様です。

// Human interface
type Human interface {
	Eat()
}

// CarlingGirl struct
type CarlingGirl struct{}

// Eat func
func (c *CarlingGirl) Eat() {
	fmt.Println("カーリング女子：もぐもぐ")
}

// FoodFighter struct
type FoodFighter struct{}

// Eat func
func (f *FoodFighter) Eat() {
	fmt.Println("フードファイター：バクバク")
}

// HumanEat func
func HumanEat(h Human) {
	fmt.Print("ご飯ですよ！ ")
	h.Eat()
}

// InterfaceFunc func
func InterfaceFunc() {
	carling := new(CarlingGirl)
	fighter := new(FoodFighter)
	HumanEat(carling)
	HumanEat(fighter)
}
