package main

import (
	"fmt"
	"math"
)

const D = 0.0000001

func sqrt(x float64) float64{
	z := 1.0
	d := 0.0

	for math.Abs(z - d) > D {
		d = z
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main(){
	const x = 5
	fmt.Println(math.Sqrt(x),math.Sqrt(x))
}