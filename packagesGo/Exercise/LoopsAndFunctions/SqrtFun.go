package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x / 2)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		if z*z == x {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9999))
	fmt.Println(math.Sqrt(9999))
}
