package main

import "fmt"

type Vehicle struct {
	maxSpeed int
}

func (v Vehicle) drive() {
	fmt.Println("vehicle drive")
}

func (v *Vehicle) setSpeed(speed int) {
	v.maxSpeed = speed
}

type Car struct {
	Vehicle
	doorsNumber int
}

func (v Car) drive() {
	fmt.Println("car drive")
}

func main() {
	car := Car{
		Vehicle: Vehicle{
			maxSpeed: 200,
		},
		doorsNumber: 5,
	}

	car.setSpeed(400)

	fmt.Println(car.maxSpeed)
}
