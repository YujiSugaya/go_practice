package main

import "fmt"

func main(){
	//vはint型と明示していないが、型推論してくれている
	v := 42
	fmt.Printf("v is of type %T\n",v)
}