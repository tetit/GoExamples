package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

func main() {
	// fmt.Println("pi", pi)

	// Заборонено призначати нові значення
	// pi = 123.2

	// const (
	// 	a = 1
	// 	b = "hello"
	// 	c = true
	// 	e = 1234.3456
	// )
	var s1 string = "hello"
	const s2 = "hello"

	// s1[0] = "q23r"
	s1 = "привіт"
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	// s2[0] = "ew3fa2"
	// s2 = "привіт"
	// fmt.Println("s2", s2)
}
