package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var i1 int = 42
	// var f1 float64 = float64(i1)
	// var u uint1 = uint1(f1)

	i1 := 42
	f1 := float64(i1)
	u := uint(f1)

	fmt.Println(i1, f1, u)

	// var x, y int = 3, 4
	x, y := 3, 7
	flo := math.Sqrt(float64(x*x + y*y))
	z := uint(flo)
	fmt.Println(x, y, z, flo)

}
