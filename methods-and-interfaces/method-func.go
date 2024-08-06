package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//メソッドとして定義するのではなく、通常の関数として定義する
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	//メソッドとして定義していないのでv.Absにはならず、通常の呼び出し
	fmt.Println(Abs(v))

}
