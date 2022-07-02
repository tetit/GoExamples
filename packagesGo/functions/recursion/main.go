package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func sumSlice(sl []int) int {
	if len(sl) == 1 {
		return sl[0]
	} else {
		i := len(sl)-1
		return sl[i] + sumSlice(sl[:i])
	}
}

func main() {

	fmt.Println(fact(5))
	sl := []int{1, 2, 3, 4, 7}
	fmt.Println(sumSlice(sl))

	
}