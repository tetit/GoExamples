package main

import "fmt"

func main() {
	// func copy(dst, src []Type) int
	// It takes in two slices dst and src, and copies data from src to dst. It returns the number of elements copied.

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	fmt.Printf("Number Of Elements Copied: %d\n", copy(dst, src))
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	dst2 := []int{6, 7, 8}
	fmt.Printf("Number Of Elements Copied: %d\n", copy(dst2, src))
	fmt.Printf("dst: %v\n", dst2)
	fmt.Printf("src: %v\n", src)

	dst3 := make([]byte, 9)
	src3 := "string some"
	fmt.Printf("Number Of Elements Copied: %d\n", copy(dst3, src3))
	fmt.Printf("dst: %v\n", dst3)
	fmt.Printf("src: %v\n", src3)

	src4 := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Number Of Elements Copied: %d\n", copy(src4, src4[3:5]))
	fmt.Printf("src4: %v\n", src4)

}
