package main

import "fmt"

func main(){
	sum := 1
// 初期値： sum = 1
// 1回目のループ： sum = 1 + 1 = 2
// 2回目のループ： sum = 2 + 2 = 4
// 3回目のループ： sum = 4 + 4 = 8
//初期化と後処理のステートメントは任意
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}