package main

import (
	"fmt"
	"math"
)

//sqrtという関数はxという引数をとり、
func sqrt(x float64) string{
	//引数が0未満の場合 -引数にiという文字を追加する
	if x < 0 {
		return sqrt(-x) + "i"
	}
	//０未満ではない場合、引数の平方根を求める
	//接頭辞にSがつく場合、フォーマットした結果を文字列で返す
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	//2は0未満ではなので、平方根を求める、-4は0未満なので4+1になる
	//次の処理が行われ2+iになる
	fmt.Println(sqrt(2),sqrt(-4))
}