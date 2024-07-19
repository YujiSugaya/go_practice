package main

import "fmt"

func main() {
	// var 変数名 [] 型名　で配列を定義する
	var a [2]string
	//変数名[インデックス番号]　で対象のインデックスを指定する
	a[0] = "hello"

	a[1] = "world"
	fmt.Println(a[0], a[1])
	//配列全部を出力する[...]
	fmt.Println(a)

	//変数の宣言と代入を同時に行う
	// 変数名 := [配列のサイズ] 型名 {中身}
	primes := [6]int{2, 3, 5, 7, 11, 13} //配列のサイズは変えることはできないよ
	fmt.Println(primes)
}
