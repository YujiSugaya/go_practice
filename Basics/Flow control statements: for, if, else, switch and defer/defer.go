package main

import "fmt"

func main() {
	//deferステートメントは defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。
	//defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。
	//順番的には最初だけど2番目に実行される
	defer fmt.Println("World")
	//deferがあるので先に実行される
	fmt.Println("hello")
}
