package main

import "fmt"

//二つ以上引数が同じ型の場合、最後に型を追加して省略できます。
//引数の後のintは戻り値の型を指定している
func add(x,y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42,13));
}