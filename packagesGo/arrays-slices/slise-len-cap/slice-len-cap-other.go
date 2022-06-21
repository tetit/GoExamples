package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[0:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[0:4]
	printSlice(s)

	s = s[0:3]
	printSlice(s)

	s = s[0:4]
	printSlice(s)

	var s1 []int
	printSlice(s1)
	if s1 == nil {
		fmt.Println("nil")
	}

	a := make([]int, 5)
	printSliceStr("a", a)

	b := make([]int, 0, 5)
	printSliceStr("b", b)

	b1 := make([]int, 3, 5)
	printSliceStr("b1", b1)

	c := b[:2]
	printSliceStr("c", c)

	d := c[2:5]
	printSliceStr("d",d)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceStr(x string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", x, len(s), cap(s), s)
}
