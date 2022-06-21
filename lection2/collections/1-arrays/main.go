package main

const size = 10

func main() {
	// var (
	// a1 [2]int //  [int int]
	// a2 [size]string
	// a3 [4][2]int // [[int int] [int int] [int int] [int int]]
	// // довжина має бути указана за допомогою цілого числа
	// // w1 [3.5]int
	// t := [...]int{1, 2, 4}
	// )

	
	// fmt.Println("a1 =", a1)
	// fmt.Println("a2 =", a2)
	// fmt.Println("a3 =", a3)

	// ініціалізація разом з даними
	// a4 := [2]int{3, 11111} // [3 11111]
	// fmt.Println("a4 =", a4)

	// ініціалізація разом з даними та порядком
	// a4 := [2]int{1: 3, 0: 11111} // [11111 3]
	// fmt.Println("a4 =", a4)

	// ініціалізація разом з даними та порядком - складніше
	// a4 := [4]int{1: 3, 3: 9} // [0 3 0 9]
	// a4 := [4]int{3, 9, 0,0} // [3 9 0 0]

	// fmt.Println("a4 =", a4)

	// вибір по індексу
	// elem0 := a4[0]
	// fmt.Println("елемент під індексом 0 =", elem0)

	// out of range - заборонено
	// elem4 := a4[len(a4)-1]
	// fmt.Println("elem4 =", elem4)

	// elem-1 := a4[-1]

	// визначення length
	// a9 := [3]string{"cogito", "ergo", "sum"}
	// fmt.Println("довжина a9 =", len(a9))

	// перезапис елементів
	// a4[0] = 25
	// fmt.Println("a4 =", a4)

	// порівняння
	// a5 := [2]int{9, 8}
	// a6 := [2]int{9, 8}
	// fmt.Println("a5 та a6 однакові? - ", a5 == a6)

	// a7 := [2]int{9, 8}
	// a8 := [3]int{9, 8}
	// fmt.Println("a7 та a8 однакові? - ", a7 == a8)
}
