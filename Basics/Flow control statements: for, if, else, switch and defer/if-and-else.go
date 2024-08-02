package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64{
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//%g: 浮動小数点数や複素数を自動的に適切な精度で表現します。
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

//27 >= 20
//9 20

// 先に 27 >= 20が出力されるのはPrintがif文の中にあるので、一番最初に実行される
// その後一行目のpowの戻り値
// その後は二行目のpowの戻り値が出力される