package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)
	//aの中には0番目から2番目まで取得する　John Paul
	a := names[0:2]
	//bの中には1番目から3番目まで取得する　Paul George
	b := names[1:3]
	fmt.Println(a, b) //[John Paul] [Paul George]
	
	//Paul GeorgeのPaulをXXXXに変更したけどa[1]もb[0]と同じものを参照しているため両方書きかわる
	b[0] = "XXX"

	fmt.Println(a, b)　//[John XXX] [XXX George]
	//元の配列の1番のデータも書き換わった。
	fmt.Println(names) //[John XXX George Ringo]
}
