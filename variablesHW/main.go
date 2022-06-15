package main

import (
	"fmt"
)

const (
	applePrice  = 5.99
	pearPrice   = 7.
	moneyAmount = 23.
)

func main() {

	appelsAmount, pearsAmount := 9, 8
	moneyToBuy := applePrice * float64(appelsAmount)
	moneyToBuy += pearPrice * float64(pearsAmount)
	fmt.Printf("1. How much money should we spend to buy %d apples and %d pears? \n", appelsAmount, pearsAmount)
	fmt.Printf("Answer is: %.2f \n", moneyToBuy)

	amountFloat := moneyAmount / pearPrice
	pearsAmount = int(amountFloat)
	fmt.Println("2. How many pears can we buy?")
	fmt.Printf("Answer is: %d \n", pearsAmount)

	amountFloat = moneyAmount / applePrice
	appelsAmount = int(amountFloat)
	fmt.Println("3. How many appels can we buy?")
	fmt.Printf("Answer is: %d \n", appelsAmount)

	appelsAmount, pearsAmount = 2, 2
	moneyToBuy = applePrice * float64(appelsAmount)
	moneyToBuy += pearPrice * float64(pearsAmount)
	canBuy := moneyAmount >= moneyToBuy
	fmt.Printf("4. Can we buy %d pears and %d apples? \n", pearsAmount, appelsAmount)
	fmt.Println("Answer is: ", canBuy)

}
