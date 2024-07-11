package main

import "fmt"
//複数の変数に対して同じ型を定義する場合、一番後ろに型を定義する
var c, python, java bool



func main(){
	//var + 変数名 +　型
	var i int 
	var a string
	//int型の初期値は0 bool型の初期値はfalse
	//string型は""
	fmt.Println(a,i,c,python,java)
}