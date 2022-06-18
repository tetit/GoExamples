package main

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	interpretedString string = "I need to escape symbols"
	rawString                = `I can use anything here "`
)

func main() {
	// firstName := "Viktor"
	// lastName := "Pakhuchy"
	// fullName := firstName + " " + lastName
	// fmt.Println("конкатинація", fullName)

	// fmt.Println("довжина рядка", len(firstName))

	// fmt.Println("порівняння", "c" == "с")

	// fmt.Println("посилання на конкретний елемент", firstName[0])

	// // заборонено
	// fmt.Println("зміна рядка")
	// firstName[0] = "s"

	// firstName = "siktor" + "a"

	// fmt.Println("правильне редагування рядків (створення нового)", strings.Replace("Hello", "l", "d", -1))
	// fmt.Println("конвертація в маленькі літери", strings.ToLower("Viktor PakHuchyi"))
	// fmt.Println("конвертація в великі літери", strings.ToUpper("Viktor Pakhuchyi"))
	// s := "String"
	// sClone := strings.Clone(s)
	// fmt.Println(&s)
	// fmt.Println(&sClone)

	// src := []int{1, 2, 3, 4, 5}
	// dst := make([]int, 5)
	// numberOfElementsCopied := copy(dst, src)

	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst: %v\n", dst)
	// fmt.Printf("src: %v\n", src)

	// //After changing dst
	// dst[0] = 10
	// fmt.Println("\nAfter changing dst")
	// fmt.Printf("dst: %v\n", dst)
	// fmt.Printf("src: %v\n", src)

	// //Length of destination is less than length of source
	// dst = make([]int, 4)
	// fmt.Printf("\n\ndst int4: %v\n", dst)
	// numberOfElementsCopied = copy(dst, src)
	// fmt.Println("\nLength of dst less than src")
	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst: %v\n", dst)
	// fmt.Printf("src: %v\n", src)

	// //Length of destination is greater than length of source
	// dst = make([]int, 6)
	// numberOfElementsCopied = copy(dst, src)
	// fmt.Println("\nLength of dst less than src")
	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst: %v\n", dst)
	// fmt.Printf("src: %v\n", src)

	// src := "abc"
	// dst := make([]byte, 3)

	// numberOfElementsCopied := copy(dst, src)
	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst: %v\n", dst)
	// fmt.Printf("src: %v\n", src)

	// dst1 := make([]byte, 5)
	// numberOfElementsCopied = copy(dst1, src)
	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst1: %v\n", dst1)

	// src2 := []string{"a", "no", "type"}
	// dst2 := make([]string, 7)
	// numberOfElementsCopied = copy(dst2, src2)
	// fmt.Printf("\n\nNumber Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("dst2: %v\n", dst2)

	// dst2 = make([]string, 2, 5)
	// numberOfElementsCopied = copy(src2, dst2)
	// fmt.Printf("\n\nNumber Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("src2: %v\n", src2)

	// numberOfElementsCopied = copy(src2, src2[2:])

	// fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	// fmt.Printf("src2: %v\n", src2)

	// s := "Separate;;../&$8e786;7 uy&6#t# text"
	// fmt.Println(strings.Cut(s, "a"))

	// show := func (s, sep string) {
	// 	before, after, found := strings.Cut(s, sep)
	// 	fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	// }
	// show("Gopher", "Go")
	// show("Gopher", "ph")
	// show("Gopher", "er")
	// show("Gopher", "Badger")

	// f := func (c rune) bool {
	// 	// fmt.Println("c = ", c, !unicode.IsLetter(c) && !unicode.IsNumber(c))
	// 	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	// }
	// arrS := strings.FieldsFunc(s, f)
	// fmt.Println(arrS)
	// fmt.Println(!unicode.IsLetter(';'))

	// b := func (c rune) bool {
	// 	return c == 'a'
	// }
	// fmt.Println(strings.FieldsFunc(s, b))

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
	fmt.Printf("%q\n", unicode.Han)

	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsLetter))

	fmt.Println("\tis mark rune")

	fmt.Println("na" + strings.Repeat("tot", 3))

	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))

	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

}
