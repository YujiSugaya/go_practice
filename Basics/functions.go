package main

import "fmt"

//関数の引数を設定するときは変数名+データ型　
//引数が複数ある場合は、で複数設定
func add(x int, y int) int {
	return x + y
}



func main(){
	fmt.Println(add(42,13))
}