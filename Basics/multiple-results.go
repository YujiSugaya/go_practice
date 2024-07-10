package main

import "fmt"

//x,yというstring型を受け取り、戻り値にstring型を設定し、受け取った文字列をy,xの順で返す
//戻り値の方は二つ以上同じ型でも省略はできない
func swap(x,y string) (string,string){
	return y,x
}

func main(){
	//代入するときは:=を使用する
	a,b := swap("Hello","World")
	fmt.Println(a,b)
}