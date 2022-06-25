package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	s1 := strings.Fields(s)
	m := make(map[string]int)
	for _, key := range s1 {
		if _, ok := m[key]; ok {
			m[key] += 1
		} else {
			m[key] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
