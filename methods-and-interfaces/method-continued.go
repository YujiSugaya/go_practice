package main

import (
	"fmt"
	"math"
)

//Myfloatという新い方を定義している　float64の機能を持ちつつ独自のメソッドを追加することができる
type MyFloat float64

//fはレシーバ変数　MyFloat型に対してAbsメソッドを定義している1
func (f MyFloat) Abs() float64 {
	if f < 0 {
		//0より小さい場合つまり負の値の場合はその値に-をかけて正の値にします。
		return float64(-f)
	}
	//正の値の場合はそのまま返します。
	return float64(f)
}

func main() {
	// - ルート2がfに入っている　-1.4142
	f := MyFloat(-math.Sqrt2)
	//fのAbsメソッドを出力する
	//今回は負の値が入っているので　正の値にして出力する
	fmt.Println(f.Abs())
}
