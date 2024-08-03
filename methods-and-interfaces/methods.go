package main 

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Vertex 型のレシーバ v に対するメソッド Abs を定義している
func (v Vertex) Abs() float64 {

	//v.Xでvという構造体のなかのxにアクセスしている vの中のXが3の場合3*3
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//vのなかの構造体を入れる
	v := Vertex{3, 4}
	//vのなかのAbsメソッドを使用する
	fmt.Println(v.Abs())
}


