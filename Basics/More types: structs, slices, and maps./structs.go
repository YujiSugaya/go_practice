package main

import "fmt"

// type 構造体名　struct {} で定義することができます。
type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
