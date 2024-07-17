package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	//vという構造の中にX＝１、Y＝２を代入する
	v := Vertex{1, 2}
	//.を使ってアクセスする
	fmt.Println(v.X) //X=1
	//vのXに４を代入する
	v.X = 4
	fmt.Println(v.X) //4
}
