package main

import (
	"fmt"
	"reflect"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

//Assign Default Value for Struct Field
func (obj *Employee) Info() {
	if obj.FirstName == "" {
		obj.FirstName = "John"
	}
	if obj.LastName == "" {
		obj.LastName = "Dou"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func main() {
	emp1 := Employee{FirstName: "Fred"}
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee{Age: 26}
	emp2.Info()
	fmt.Println(emp2)

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

	fmt.Println(e.EmpInfo())

	fmt.Println(reflect.TypeOf(emp1))
	fmt.Println(reflect.ValueOf(e).Kind())
}
