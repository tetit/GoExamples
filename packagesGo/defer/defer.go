package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 47
}

func main() {
	defer fmt.Printf("\ndefer")
	fmt.Print("Hello I'm ")
	fmt.Print("counting ")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")

	}

	fmt.Println("done")

	fmt.Println(c())

}
