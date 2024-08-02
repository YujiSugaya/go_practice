package main

import "fmt"


func main(){
	sum := 1
	//;は省略することができる
	//whileはgolangにはなくforだけ使う
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)
}