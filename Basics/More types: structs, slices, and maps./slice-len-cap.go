package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]     //0番目から0番を取得するので長さは0でも６つ値は入るよ　データは保持されたまま
	printSlice(s) //len=0 cap=6 []

	fmt.Println(s)
	//スライスの長さを0から4に拡張する　もともと入っていたデータが参照されるので0ではなく、2357
	s = s[:4] // 2357を入れる
	printSlice(s) //len=4 cap=6 [2 3 5 7]

	//5 7 
	s = s[2:]
	printSlice(s)　//len=2 cap=4 [5 7]
	//2: でインデックスが0番、一番のデータを削除しているため、capは6-2で4になる
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//スライスは長さ( length )と容量( capacity )の両方を持っています。
//スライスの長さは、それに含まれる要素の数です。
//スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数です。
