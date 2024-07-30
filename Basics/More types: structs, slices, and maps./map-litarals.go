package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	//キー "Bell Labs" に対応するVertex構造体（緯度 40.68433, 経度 -74.39967）
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	//キー "Google" に対応するVertex構造体（緯度 37.42202, 経度 -122.08408）
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
