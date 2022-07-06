package main

import "fmt"

func main() {
	// var i *int
	// *i = 1

	// var n []int
	// _ = n[10]

	// panic("Help mi!")
	someFun1()
}

func someFun1() (n int) {
	defer func() {
		recover()
		fmt.Println("defer 1")
		n = 2
	}()
	fmt.Println("start 1")
	someFun2()
	fmt.Println("end 1")
	defer func() {
		fmt.Println("defer 2")
	}()
	return 1
}
func someFun2() {
	defer func() {
		recover()
		fmt.Println("defer 3")
	}()
	fmt.Println("start 2")
	panic("help me!")
	fmt.Println("end 2")

}
