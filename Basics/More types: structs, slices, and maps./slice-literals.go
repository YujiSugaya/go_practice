package main

import "fmt"

func main() {
	//intのスライス
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	//boolのスライス
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	//structのスライス

	s := []struct {
		//iがint bがbool
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Println(s[0].i)
}