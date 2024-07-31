package main

import (
	"fmt"
	"math"
)

//compute 関数は、2つの float64 を引数に取り float64 を返す関数を引数として受け取り、その関数に3と4を渡して結果を返します。

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	//hypot は匿名関数（無名関数）として定義されています

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) //169の平方根なので13が出力される

	fmt.Println(compute(hypot))// 9+16= 25 → 5
	fmt.Println(compute(math.Pow)) //3 x 3 x 3 81
}
