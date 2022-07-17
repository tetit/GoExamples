package main

import (
	"GoExamples/lection8/mobile"
	"fmt"
)

func main() {
	fmt.Println(mobile.ParsePhoneNumber("+38012333345", "UA"))
	fmt.Println(mobile.ParsePhoneNumber("+89", "UK"))
}