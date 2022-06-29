package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "世界"

	bb := []byte(s) // byte = uint8
	rr := []rune(s) // rune = int32

	fmt.Println("bytes:", bb)
	fmt.Println("runes:", rr)

	fmt.Println()
	for i, b := range bb {
		fmt.Println(i, "byte:", string(b))
	}

	fmt.Println()
	for i, r := range rr {
		fmt.Println(i, "rune:", string(r))
	}

	fmt.Println(split(s + " Jon Smit Yun"))
}

// Написати функцію, що приймає рядок повного імені, де частини імен розділені пробілами
//(example: “Harry James Potter”), повернути слайс частин імен.
func split(fullName string) []string {
	const nameSeparator = " "

	names := strings.Split(fullName, nameSeparator)

	return names
}
