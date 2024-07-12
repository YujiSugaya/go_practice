package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	//1<<64 -1　　　
	// << は左ビットシフト演算子
	// 元の数値を指定されたビット数だけ左に日どうさせるその後0ビットを右端に追加する
	// 1<<64
	// 1というのは64ビット中1
	// その1を64ビットに左シフトする uint64型では64ビット以上のシフト結果は0になる
	// 左にシフトした後-1するという意味

	//18446744073709551615
	MaxInt uint64 = 1<<64 -1



	//複素数の平方根を計算する関数
	// z = -5 + 12i
	// a = -5, b = 12
	// √(-5)の二乗+12の二乗=√169 = 13
	// 実部 √2分の13+(-5) =√2分の8 =√4 = 2
	// 虚部 sign(12)・√2分の13-(-5) = 3
	// したがって２ + 3i


	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Println(MaxInt)
	fmt.Println(z)
	//%Tは変数の方を出力する
	//%vは変数の値のデフォルトを出力する
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", z, z)
	fmt.Printf("Type: %T value: %v\n", z, z)
}
