package main

import "fmt"

func main() {
	a := make([]int, 5)
	//すべての要素が0で初期化されています。
	printSlice("a", a) //a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	//長さ0、容量5の整数スライスを作成します。
	printSlice("b", b) //b len=0 cap=5 []

	//bから新しいスライスcを作成します。出力: c len=2 cap=5 [0 0]
	c := b[:2]
	//bの容量全体（5）を引き継ぎますが、長さは2になります。
	//bは空でしたが、基底配列の最初の2要素（0で初期化）を参照します。
	printSlice("c", c) //c len=2 cap=5 [0 0]

	d := c[2:5]
	//cから新しいスライスdを作成します。
	//出力: d len=3 cap=3 [0 0 0]
	//cの2番目のインデックスから始まり、5番目のインデックス（排他的）まで選択します。
	printSlice("d", d) //d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
