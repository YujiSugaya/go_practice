package main

import "fmt"

func main() {
	//10個の要素を持つスライスを作成
	pow := make([]int, 10)

	for i := range pow {
		//iにr対して1 << uint(i)をp計算する
		//左シフト演算子を使って2のi乗を計算する
		pow[i] = 1 << uint(i) // == 2**i
	}
	_// はインデックスを無視することを示します
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
