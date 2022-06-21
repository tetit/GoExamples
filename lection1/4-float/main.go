package main

import (
	"fmt"
	"math"
)

var (
	f1 float32 = 1
	f2 float64 = 254
	f3 float64 = 25.0011
	f4         = -13.9098
	f5         = 13.9
	f6         = 13
	f7 float64 = 13
)

func main() {
	// fmt.Println("дано", f3, f4)
	// mult := f3 * f4
	// fmt.Println("множення", mult)

	// someInt := 10
	// fmt.Println("дано", f3, someInt)
	// mult := int(f3) * someInt
	// fmt.Println("хибне множення", mult)

	// someInt := 10
	// fmt.Println("дано", f3, someInt)
	// var mult float64 = f3 * float64(someInt)
	// fmt.Println("правильне множення", mult)

	// fmt.Println("максимальне значення f32", math.MaxFloat32)
	// fmt.Println("максимальне значення f64", math.MaxFloat64)
	fmt.Println(int(math.Abs(f4)))
	fmt.Println(int(math.Acos(f4)))
	fmt.Println(int(math.FMA(2.1, 2.5, 2.2)))
	fmt.Println(int(math.Float32bits(f1)))
	fmt.Println(math.Floor(f4), math.Floor(f5))
	fmt.Println(math.Trunc(f4), math.Trunc(f5))

}
