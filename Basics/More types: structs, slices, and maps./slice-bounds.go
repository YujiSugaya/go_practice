package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11} 

	s = s[1:4] //再代入
	fmt.Println(s) //[3 5 7]

	s = s[:2]　// == [0:2] 先頭がない場合はデフォルトで0になる
	fmt.Println(s) // 3 5

	s = s[1:] // == [1:max] 5　後ろがない場合はデフォルトでマックスになる
	fmt.Println(s) //5
}
