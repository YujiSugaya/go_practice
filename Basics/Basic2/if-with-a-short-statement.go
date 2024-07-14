package main

import (
	"fmt"
	"math"
)
//pow 関数は3つの引数を受け取ります：x（基数）、n（指数）、および lim（制限値）。
//if文の条件の中で代入を行い判定する。　
func pow(x, n, lim float64) float64{
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	//ifのスコープ外なのでvは使用できない
	// return v
}


func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}