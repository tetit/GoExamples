package main

import "fmt"

var c, python, java bool
var j, l int = 1, 2

func main()  {
	var i int
	fmt.Println(i, c, python, java)
	var c = true
	fmt.Println(c)
	c = false	
	fmt.Println(c)

	var c1, python1, java1 = true, false, "no!"
	fmt.Println(c1, python1, java1, j, l)
	fmt.Printf("\n%T", c1)
	fmt.Printf("\n%T", python1)
	fmt.Printf("\n%T", java1)

	q, w, e, r := "text", 5.007, 48, true
	fmt.Println(q, w, e, r) //?
	fmt.Println(q)

}