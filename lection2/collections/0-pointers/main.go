package main

import "fmt"

const name = "Boris"

var age int = 20

func main() {
	// fmt.Println("name=", name)
	// fmt.Println("age=", age)

	// поінтер на age
	// fmt.Println("&age=", &age)

	var pAge *int
	fmt.Println("pAge=", pAge)
	// pAge++
	n := 1 
	pAge = &n
	fmt.Println("pAge=", *pAge+1)

	// що таке nil
	// pAge = nil

	pAge = new(int)
	fmt.Println("занчення на яке вказує змінна *pAge=", *pAge)
	fmt.Println("значення самої змінної pAge=", pAge)

	pAge = &age
	fmt.Println("значення самої змінної pAge=", pAge)
	fmt.Println("занчення на яке тепер вказує змінна *pAge (а вказує вона на значення змінної age)=", *pAge)
	// змінюємо значення age
	age++
	fmt.Println("значення змінної age=", age)
	fmt.Println("занчення на яке вказує змінна *pAge=", *pAge)
	fmt.Println("занчення pAge=", pAge)

	*pAge = 99
	fmt.Println("значення змінної age=", age)
	fmt.Println("занчення на яке вказує змінна *pAge=", *pAge)

	// поінтер на поінтер
	// original := "what the f..."
	// p := &original
	// fmt.Println("значення, на котре вказує змінна p=", *p)
	// fmt.Println("значення змінної p=", p)
	// fmt.Println("--------")
	// pp := &p
	// fmt.Println("значення, на котре вказує змінна pp=", *pp)
	// fmt.Println("значення змінної pp=", pp)
	// fmt.Println("оригінальне значення=", **pp)

	// fmt.Println("-----")

	// d := 9
	// fmt.Println("значення змінної d=", d)
	// l := d
	// fmt.Println("значення змінної l=", l)
	// d++
	// fmt.Println("значення змінної d=", d)
	// fmt.Println("значення змінної l=", l)

	// порівняння поінтерів
	// var v1 int = 111
	// var v2 int = 111
	// fmt.Println("значення змінної v1=", v1)
	// fmt.Println("значення змінної v2=", v2)
	// fmt.Println("поінтер на v1=", &v1)
	// fmt.Println("поінтер на v2=", &v2)
	// fmt.Println("поінтер на v1 та v2 однакові?", &v1 == &v2)

	// *nil - заборонено
	// var b *string
	// fmt.Println(b)

	// if b != nil{
	// 	fmt.Println(*b)
	// }

	// & - щоб взяти поінтер на щось
	// * - шоб взяти значення на яке вказує поінтер

	// поінтер на константу - заборонено
	// fmt.Println("&name=", &name)
}
