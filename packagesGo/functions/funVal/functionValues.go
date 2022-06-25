package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {

	fmt.Println("in compute")
	return fn(3, 4)
}

func adder() func(int) int {
	fmt.Println("in adder")
	sum := 0
	return func(x int) int {
		fmt.Println("in func")
		sum += x
		return sum
	}
}

func newCounter() func() int {
	GFG := 0
	return func() int {
		GFG += 1
		return GFG
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		fmt.Println("in hypot")
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	GFG := 0

	counter := func() int {
		GFG += 1
		return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())

	counter = newCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	anotherCounter := newCounter()
	fmt.Println(anotherCounter())
	fmt.Println(anotherCounter())
	fmt.Println(anotherCounter())

}
