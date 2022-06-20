package main

import (
	"fmt"
	"strings"
)

const size = 10

func main() {
	// var s1 [][]bool // дефолтне значення слайсу - nil; cap=0 len=0 array=nil
	// fmt.Println("s1", s1)

	// слайс слайсів
	// var s2 = [][]string{
	// 	{"A", "B", "C"},
	// 	{"D", "E", "F"},
	// }
	// fmt.Println("s2", s2)

	// довжина та капасіті
	// fmt.Println("s2 length =", len(s2))
	// fmt.Println("s2 capacity =", cap(s2))

	// чи слайс пустий
	// s3 := []int{}
	// var s4 []int
	// fmt.Println("чи слайс пустий? - s3", s3 == nil)
	// fmt.Println("чи слайс пустий? - s4", s4 == nil)

	// fmt.Println("чи слайс пустий? - s3", len(s3))
	// fmt.Println("чи слайс пустий? - s4", len(s4))

	// fmt.Println("чи слайс пустий? - ", s5 == nil)
	// fmt.Println("чи слайс пустий? - ", len(s4))

	// запис в пустий/nil слайс - заборонено
	// s3[0] = 5
	// s4[0] = 5

	// створення через make
	// s4 := make([]int, 2, 4)
	// fmt.Println("s4 =", s4)
	// fmt.Println("s4 length =", len(s4))
	// fmt.Println("s4 capacity =", cap(s4))

	// створення через зріз array
	// s5 := s4[:1]
	// fmt.Println("s5 =", s5)
	// fmt.Println("s5 length =", len(s5))
	// fmt.Println("s5 capacity =", cap(s5))

	// s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s6 := s5[2:4] // [start : stop]
	// fmt.Println("s6 =", s6)
	// fmt.Println("s6 length =", len(s6))
	// fmt.Println("s6 capacity =", cap(s6))
	// fmt.Println("s6 capacity =", cap(s6))

	// s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s6 := s5[:]
	// s7 := s6[1:]
	// fmt.Println("s5", s5)
	// fmt.Println("s6", s6)
	// fmt.Println("s7", s7)

	// s7[3] = 99
	// fmt.Println("-----")
	// s5[8] = 456
	// fmt.Println("s5", s5)
	// fmt.Println("s6", s6)
	// fmt.Println("s7", s7)

	// append
	s7 := make([]int, 2)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s7 =", s7)
	fmt.Println("s7 length =", len(s7))
	fmt.Println("s7 capacity =", cap(s7))

	s8 := append(s7, 99, 88)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s8 =", s8)
	fmt.Println("s8 length =", len(s8))
	fmt.Println("s8 capacity =", cap(s8))

	s9 := append(s7, s8...)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s9 =", s9)
	fmt.Println("s9 length =", len(s9))
	fmt.Println("s9 capacity =", cap(s9))

	s9 = append(s9, 101)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s9 =", s9)
	fmt.Println("s9 length =", len(s9))
	fmt.Println("s9 capacity =", cap(s9))

	s9 = append(s9, s8...)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s9 =", s9)
	fmt.Println("s9 length =", len(s9))
	fmt.Println("s9 capacity =", cap(s9))

	s9 = append(s9[:4], s9[7:]... )
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("s9 =", s9)
	fmt.Println("s9 length =", len(s9))
	fmt.Println("s9 capacity =", cap(s9))

	s10 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(s10[1:3], append(s10[:], 6, 7))
	fmt.Println(s10)
	fmt.Printf("%T\n", s10)
	s11 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", s11)
	fmt.Println(append(s11, s10[:]...))

	for idx, s := range s11 {
		fmt.Println(idx, s)
	}

	for i := 0; i < len(s10); i++ {
		fmt.Println(s10[i])
	}


    words := []string{}
    words = append(words, "an")
    words = append(words, "old")
    words = append(words, "falcon")

    res := strings.Join(words, " ")
	fmt.Println(words)
    fmt.Println(res)

	s12 := make([]int, 7, 15)
	fmt.Println(s12)
	
	s13 := make(map[int]bool)
	fmt.Println(s13)

}
