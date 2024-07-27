package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//mapの初期化
	m = make(map[string]Vertex)
	//キー"Bell Labs"に対応する値として新しいVertex構造体を作成し、マップに追加しています。
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
