package main

import "fmt"

func main() {
	//i,jにそれぞれの数値を代入する
	i, j := 42, 2701

	//pはポイント変数
	//&i　はi(42)のメモリアドレスを取得する演算子
	//pを初期化している
	p := &i
	//0xc00001e0b8がpの中に入っている
	// fmt.Println(p)

	//*はポインタリファレンスの演算子
	//pが指しているメモリの位置の値にアクセスるする
	fmt.Println(*p) //42
	//ポインタpが指しているメモリ位置にある値を21に設定します。
	*p = 21
	//iの中身は42
	fmt.Println(i)

	//pの中にj(2701)のメモリアドレスを取得して代入する
	p = &j
	// fmt.Println("aa")
	// fmt.Println(p) 0xc00001e0c0
	// fmt.Println(*p) 2071
	//pのポインタメモリ位置の値を現在の＊p(21)/37 して、その結果を*pに代入します
	*p = *p / 37
	//jの値を出力する　2701/37 =73
	//*p = の段階でjの値を73に更新している
	// fmt.Println(*p)
	fmt.Println(j)
}
