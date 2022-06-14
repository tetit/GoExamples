package main

import (
	"fmt"
)

var (
	applePrice  = 5.99
	pearPrice   = 7.
	moneyAmount = 23.
)

func main() {
	var (
		appelsAmount int
		pearsAmount  int
		moneyToBuy   float64
	)

	appelsAmount, pearsAmount = 9, 8
	moneyToBuy = applePrice*float64(appelsAmount) + pearPrice*float64(pearsAmount)
	fmt.Printf("1. How much money should we spend to buy %v apples and %v pears? \n", appelsAmount, pearsAmount)
	fmt.Printf("Answer is: %v\n", moneyToBuy)

	pearsAmount = int(moneyAmount / pearPrice)
	fmt.Println("2. How many pears can we buy?")
	fmt.Println("Answer is: ", pearsAmount)

	appelsAmount = int(moneyAmount / applePrice)
	fmt.Println("3. How many appels can we buy?")
	fmt.Println("Answer is: ", appelsAmount)

	appelsAmount, pearsAmount = 2, 2
	moneyToBuy = applePrice*float64(appelsAmount) + pearPrice*float64(pearsAmount)
	canBuy := moneyAmount >= moneyToBuy
	fmt.Printf("4. Can we buy %v pears and %v apples? \n", pearsAmount, appelsAmount)
	fmt.Println("Answer is: ", canBuy)

}
