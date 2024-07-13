package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	//1 << 100 は、1 を 100 ビット左にシフトします。これは 1 に続いて 100 個のゼロビットが続くことを意味します。
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	//非常に大きな整数 Big を 99 ビット右にシフトすることを意味します。
	Small = Big >> 99
)

func needInt(x int) int {return x * 10 +1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	// fmt.Println(Big)
	// fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

