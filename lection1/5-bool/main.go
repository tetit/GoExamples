package main

import "fmt"

var b1 bool = false
var b2 bool = true

func main() {
	// lessonsStarted := true
	// fmt.Println("чи почалось заняття?", lessonsStarted)
	// fmt.Println("чи заняття не почалось?", !lessonsStarted)

	// fmt.Println("чи 1 дорівнює 0?", 1 == 0)

	milkPrice := 33.99
	beerPrice := 25.44

	// fmt.Println("дано", milkPrice, beerPrice)
	// milkCheaper := milkPrice < beerPrice
	// fmt.Println("чи молоко дешевше?", milkCheaper)

	// fmt.Println("дано", milkPrice, beerPrice)
	// isPriceEqual := milkPrice == beerPrice
	// fmt.Println("чи ціна однакова?", isPriceEqual)

	// myMoney := 60
	// fmt.Println("дано", milkPrice, beerPrice)
	// canBuyBothThings := (milkPrice + beerPrice) < float64(myMoney)
	// fmt.Println("чи можу я придбати обидва товари?", canBuyBothThings)

	carPrice := 500000.99
	// fmt.Println("дано", milkPrice, beerPrice, carPrice)
	// carMoreExpensive := carPrice > milkPrice && carPrice > beerPrice
	// fmt.Println("чи авто дорожче двох інших товарів?", carMoreExpensive)

	fmt.Println("дано", milkPrice, beerPrice, carPrice)
	// m := milkPrice > carPrice || milkPrice > beerPrice
	// fmt.Println("чи молоко дешевше хоча б одного товару?", m)

	// fmt.Println("1", milkPrice > carPrice)
	// fmt.Println("2", milkPrice > beerPrice)
	// fmt.Println("|| - чи", milkPrice > carPrice || milkPrice > beerPrice)
	milkMoreExpensiveThanCar := milkPrice > carPrice
	milkMoreExpensiveThanBeer := milkPrice > beerPrice
	fmt.Println("&& - та/і", milkMoreExpensiveThanCar && milkMoreExpensiveThanBeer)
	fmt.Println("|| - та", false && false)
}
