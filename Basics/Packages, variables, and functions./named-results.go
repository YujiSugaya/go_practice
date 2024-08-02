package main

import "fmt"

//変数にx,yという名前をつけて実施
func split(sum int)(x,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Println(split(500))
}