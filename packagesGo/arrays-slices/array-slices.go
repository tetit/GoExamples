package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Array"
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b = [2]string{"Hello", "new array"}
	fmt.Println(b)
	fmt.Println(b[1])

	var s []int = primes[1:4]
	s1 := primes[0:6]
	fmt.Println(s)
	fmt.Println(s1)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	n1 := names[0:2]
	n2 := names[1:3]
	fmt.Println(n1, n2)

	n2[0] = "XXX"
	fmt.Println(n1, n2)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:]
	fmt.Println(s3)

	s3 = s3[0:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)
}
