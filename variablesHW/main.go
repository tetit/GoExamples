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

	appelsAmount, pearsAmount := 9., 8.
	moneyToBuy := applePrice * appelsAmount
	moneyToBuy += pearPrice * pearsAmount
	fmt.Printf("1. How much money should we spend to buy %d apples and %d pears? \n", int(appelsAmount), int(pearsAmount))
	fmt.Printf("Answer is: %.2f \n", moneyToBuy)

	pearsAmount = moneyAmount / pearPrice
	fmt.Println("2. How many pears can we buy?")
	fmt.Printf("Answer is: %d \n", int(pearsAmount))

	appelsAmount = moneyAmount / applePrice
	fmt.Println("3. How many appels can we buy?")
	fmt.Printf("Answer is: %d \n", int(appelsAmount))

	appelsAmount, pearsAmount = 2, 2
	moneyToBuy = applePrice * appelsAmount
	moneyToBuy += pearPrice * pearsAmount
	canBuy := moneyAmount >= moneyToBuy
	fmt.Printf("4. Can we buy %d pears and %d apples? \n", int(pearsAmount), int(appelsAmount))
	fmt.Println("Answer is: ", canBuy)

}
