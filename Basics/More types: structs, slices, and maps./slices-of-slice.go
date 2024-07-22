package main

import (
	"fmt"
	"strings"
)

func main() {
	//ボードの作成 3x3の二次元スライスを作成し、すべてのマスを"_"で初期化する
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	//一行目の一番目
	board[0][0] = "X"
	//三行目の3番目
	board[2][2] = "O"
	//二行目の3番目
	board[1][2] = "X"
	//二行目の一番目
	board[1][0] = "O"
	//一行目の3番目
	board[0][2] = "X"

	//ボードの各行をループで処理します。
	for i := 0; i < len(board); i++ {
		//各行の要素を空白で区切って1つの文字列にします。
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
