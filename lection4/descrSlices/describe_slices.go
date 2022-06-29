package main

import (
	"fmt"
	"math/rand"
)

// Задача: згенерувати N слайсів по 15 випадкових
// цілих чисел (не більше 10 кожне),
// по кожному слайсу вивести суму, мінімум, максимум і моду.
const (
	elemsInSlice = 15
	maxNumber    = 10
)

func describeSlices(n int) {
	if n <= 0 {
		return
	}

	slices := generateSlices(n, elemsInSlice, maxNumber)

	// describe each slice
	for i, slice := range slices {
		sum, min, max, mode, modeMet := describeSlice(slice)

		fmt.Println(fmt.Sprintf("slice #%d:", i+1), slice)
		fmt.Println("sum:", sum)
		fmt.Println("min:", min)
		fmt.Println("max:", max)
		fmt.Println("mode:", mode, "met:", modeMet)
		fmt.Println()
	}
}

func generateSlices(n, elemsInSlice, maxNumber int) [][]int {
	slices := make([][]int, n)

	for i := 0; i < n; i++ {
		slice := make([]int, elemsInSlice)

		for j := 0; j < elemsInSlice; j++ {
			slice[j] = rand.Intn(maxNumber)
		}

		slices[i] = slice
	}

	return slices
}

func describeSlice(slice []int) (sum, min, max, mode, modeMet int) {
	min = maxNumber

	entriesCount := make(map[int]int)

	for _, elem := range slice {
		sum += elem

		if elem < min {
			min = elem
		}

		if elem > max {
			max = elem
		}

		entriesCount[elem] += 1
	}

	for elem, count := range entriesCount {
		if count > modeMet {
			mode = elem
			modeMet = count
		}
	}

	return sum, min, max, mode, modeMet
}

func describeSliceNotRefactored(n int) {
	if n <= 0 {
		return
	}

	slices := make([][]int, n)

	// generate slices
	for i := 0; i < n; i++ {
		slice := make([]int, elemsInSlice)

		for j := 0; j < elemsInSlice; j++ {
			slice[j] = rand.Intn(maxNumber)
		}

		slices[i] = slice
	}

	// describe each slice
	for i, slice := range slices {
		var (
			sum int
			min = maxNumber
			max int
		)

		entriesCount := make(map[int]int)

		for _, elem := range slice {
			sum += elem

			if elem < min {
				min = elem
			}

			if elem > max {
				max = elem
			}

			entriesCount[elem] += 1
		}

		var mode, modeMet int

		for elem, count := range entriesCount {
			if count > modeMet {
				mode = elem
				modeMet = count
			}
		}

		fmt.Println(fmt.Sprintf("slice #%d:", i), slice)
		fmt.Println("sum:", sum)
		fmt.Println("min:", min)
		fmt.Println("max:", max)
		fmt.Println("mode:", mode, "met:", modeMet)
		fmt.Println()
	}
}

func main() {
	describeSlices(2)
}
