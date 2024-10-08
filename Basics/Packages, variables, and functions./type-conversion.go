package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4
	//9+16=25 √25 =5 float64なので5.0
	var f float64 = math.Sqrt(float64(x * x + y * y))
	// uint型に変更するため 5になる
	var z uint = uint(f)
	fmt.Println(x, y, z)
}


// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// 型名 +型の違う変数名