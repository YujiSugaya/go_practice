package main

import "fmt"

func main() {
	//上から順番なので一番最初に実行される
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		//0-9まで出力される予定だがdeferステートメントあり
		//deferを使っているので後ろから順番に出力される 98765453210
		defer fmt.Println(i)
	}

	//2番目に実行される
	fmt.Println("done")
}
