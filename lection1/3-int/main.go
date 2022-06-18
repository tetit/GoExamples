package main

import "fmt"

// Різниця між int8 та uint8 та byte
var (
	i1 int8  = 124
	u  uint8 = 254
	b  byte  = 124
	c int = 124
)

// int16
var i2 int16 = 32142

// Звичайний int
var i3 int = 1000000

// Розділ знаків
var i4 int = 1_000_000

func main() {
	// fmt.Println("без розділу знаків", i3)
	// fmt.Println("з розділом знаків", i4)

	// // int та byte
	// _ = i1 == b

	// fmt.Println("дано", i1)
	// i1 = i1 + 1
	// fmt.Println("додавання", i1)

	// fmt.Println("дано", i1)
	// i1 += 1
	// fmt.Println("коротка форма додавання", i1)

	// fmt.Println("дано", i1)
	// fmt.Println("ділення", i1/2)

	// fmt.Println("дано", i1)
	// i1++
	// fmt.Println("інкремент", i1)

	// fmt.Println("дано", i1)
	// i1--
	// fmt.Println("декремент", i1)

	// fmt.Println("дано", i1, i2)
	// sum := int16(i1) + i2
	// fmt.Println("сума", sum)

	println()
	fmt.Println("дано", i1, u)
	var overflow int = 80 + 80
	var a int = 1
	var b int = 1000000
	var c  = a + b

_= c
	fmt.Println("переповнення", overflow)
}
