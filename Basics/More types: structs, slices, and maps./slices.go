package main

import "fmt"

func main() {
	//これ長さを指定してるので配列
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//最初の数字はスタートする位置のインデックス、後ろの数字はどこの数字まで取得するかのインデックスじゃない位置
	//これは長さを指定していないのでスライス
	var s []int = primes[1:4]
	fmt.Println(s)
}
