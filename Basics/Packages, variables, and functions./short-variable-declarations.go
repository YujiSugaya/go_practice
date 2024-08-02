package main

import "fmt"

func main(){
	var i, j int = 1,2
	//var宣言の代わりに :=の代入文を使い、暗黙的な宣言をすることができる
	//関数の外ではvarやfuncが必要で、:=の暗黙的な宣言は利用できない
	k := 3
	c, python, java := true, false, "no!"


	//1 2 3 true false no!
	fmt.Println(i, j, k,  c, python, java)
}

//メモ　https://qiita.com/Winesburg_Ohio/items/16ef144b48b241241ecd
// :=
// ・varとタイプ型を省略（自動判断してくれるのか、初期化の時間・容量に無駄がでるのか）
// ・関数内のみで宣言可能

// var
// ・指定したタイプ型によって初期化される（ポインタ型はnilで初期化）
// ・関数外でも宣言可能