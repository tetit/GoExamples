package main

import (
	"fmt"
)

type Rectangle struct {
	a, b int
}

type square struct {
	a int
}

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	nick      string
	age       int
}

func (r Rectangle) printRactangle(position string) {
	if r.b > r.a {
		r.b, r.a = r.a, r.b
	}
	if r.a > r.b {
		if position == "horizontal" {
			printAB(r.a, r.b)
		} else if position == "vertical" {
			printAB(r.b, r.a)
		} else {
			println("wrong position")
		}
	}
}

func (r *Rectangle) resizing(n uint) {
	r.a *= int(n)
	r.b *= int(n)
}

func (r Rectangle) compareArea(rC Rectangle) {
	if r.a*r.b == rC.a*rC.b {
		fmt.Println("equal Square")
	} else {
		fmt.Println("not equal Square")

	}
}

func (r Rectangle) squareInRectangle(s square) int {
	return r.a * r.b / (s.a * s.a)
}

func printAB(a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Print("O")
		}
		fmt.Print("\n")
	}
}
func printUser(users []User) {
	for _, v := range users {
		if v.firstName == "John" && v.age > 20 {
			fmt.Println(v.firstName, v.lastName, v.nick)
		}

	}
}

func main() {
	rect := Rectangle{3, 5}
	rect.printRactangle("vertical")
	rect.resizing(2)
	rect.printRactangle("vertical")
	fmt.Println(rect)
	rect2 := new(Rectangle)
	rect2.a = 8
	rect2.b = 10
	var rect3 = new(Rectangle)
	rect3.a = 1
	rect3.b =2
	fmt.Println(rect2)
	fmt.Println(rect3)
	rect.compareArea(*rect2)
	fmt.Println(rect.squareInRectangle(square{5}))

	users := []User{
		User{
			id:        "p0ZoB1FwH6",
			email:     "john.doe@example.com",
			firstName: "John",
			lastName:  "Doe",
			nick:      "johnnydoe",
			age:       42,
		},
		User{
			id:        "p0ZoB1FwK6",
			email:     "john.mizh@example.com",
			firstName: "John",
			lastName:  "Mizh",
			nick:      "johnnymizh",
			age:       15,
		},
		User{
			id:        "p0MoB1FwK6",
			email:     "anna.karol@example.com",
			firstName: "Anna",
			lastName:  "Karol",
			nick:      "annakarol",
			age:       19,
		},
	}

	printUser(users)

	anna := User{
		id:        "p0MoB1FwK6",
		email:     "anna.karol@example.com",
		firstName: "Anna",
		lastName:  "Karol",
		nick:      "annakarol",
		age:       19,
	}
	john := User{
		id:        "p0ZoB1FwK6",
		email:     "john.mizh@example.com",
		firstName: "John",
		lastName:  "Mizh",
		nick:      "johnnymizh",
		age:       15,
	}
	someUsers := []User{anna, john}
	fmt.Println(users)
	fmt.Println(someUsers)
	

}
