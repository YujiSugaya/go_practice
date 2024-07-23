package main

import "fmt"

// powに2の累乗の整数スライスを定義する
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//rangeはスライスや配列を簡単にループさせることができる　for index, value := range collection { ... }
	for i, v := range pow {
		//iはスライスのインデックス番号、vはスライスの中身
		//なので一回目は0,1 2回目は1,2のように順番に処理が行われる
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
