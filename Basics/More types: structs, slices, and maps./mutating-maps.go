package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])//42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])//48に更新される

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])//48が削除され0になる

	v, ok := m["Answer"]
	// m に key があれば、変数 ok は true となり、存在しなければ、 ok は false となります。
	//vの中には削除されたanswerが代入され、変数が存在しないため、okという変数はfalseになる
	fmt.Println("The value:", v, "Present?", ok)
}
