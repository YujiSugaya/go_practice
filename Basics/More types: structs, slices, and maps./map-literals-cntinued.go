package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// 省略しないで書くと
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{Lat: 40.68433, Long: -74.39967},
// 	"Google":    Vertex{Lat: 37.42202, Long: -122.08408},
// }




func main() {
	fmt.Println(m)
}
