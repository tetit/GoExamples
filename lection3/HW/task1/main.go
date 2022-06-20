package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	mapResult := make(map[int]bool)

	for _, key := range arr {
		if _, value := mapResult[key]; !value {
			mapResult[key] = true
			result = append(result, key)
		}
	}

	fmt.Println(result)

	
}
