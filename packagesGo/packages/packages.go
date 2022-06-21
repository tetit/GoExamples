package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(100)
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Println(rand.Float64())

	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is ", rand.Intn(100))
}
