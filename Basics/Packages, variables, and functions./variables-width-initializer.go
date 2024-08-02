package main

import "fmt"

// var宣言では、変数ごとに初期化(initializer)を与えることができる
//初期化子が与えられている場合、型を省略することができる
var i, j int = 1, 2

func main(){
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//出力結果　1 2 true false no!