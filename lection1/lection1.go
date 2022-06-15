package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Слава Україні!")
	println("Героям слава!")
	println("Слава Україні!")
	println(`Героям слава!%"`)

	s := "Some text"
	s1 := strings.ToLower(s)
	fmt.Println(s1)

	var i int = 1000000
	var i1 int = 1_000_000

	println(i, i1)
	i++
	// println(i++)
	println(i)
}
