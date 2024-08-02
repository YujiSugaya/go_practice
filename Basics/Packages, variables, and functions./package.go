package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
}

// https://pkg.go.dev/math/rand
//randパッケージの関数で引数に整数の範囲を指定する
//0 - 引数-1
//10の場合は 10-1なので0-9の値がランダムで出力されます。
