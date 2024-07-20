package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	//&v によって、v のメモリアドレスを指すようになります
	p := &v

	//つまりv.X
	p.X = 1e9

	//Xは1e9に更新される Yは初期値
	fmt.Println(v) //{1000000000 2}

}
