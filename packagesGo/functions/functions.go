package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x

}

func print(a, b string) {
	fmt.Println(a, b)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - y
	return
}

func main() {
	fmt.Println(add(42, 13))
	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Println(sum)
	fmt.Println(add1(402, 13))
	a, b := swap("hello", "Teti")
	fmt.Println(a, b)
	fmt.Println(swap("happy now", "I am"))
	print("Funcion", "without return")
	s, s1 := split(sum)
	fmt.Println(s, s1, "naked return")
	fmt.Println(split(46))

}
