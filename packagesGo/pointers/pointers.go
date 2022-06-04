package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	fmt.Println(i, j)

	var p *int = &i // p := &i same
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
	fmt.Println(i)
	
	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(*p)
	fmt.Println(j)



	
}