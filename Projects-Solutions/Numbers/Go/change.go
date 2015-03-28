package main

import (
	"fmt"
	"math"
	"os"
)

//=====Change Return Program=======
//The user enters a cost and then the amount of money given.
//The program will figure out the change and the number of quarters, dimes, nickels, pennies needed for the change.

//Finds the number of bills or coins to return e.g. how many bills of $100 should be returned ?
func balance(change, value float64) float64 {
	count := float64(0)

	if value <= change {
		count = math.Floor(change / value)
	}

	return count
}

//Display output to the screen
func display(number float64, s string) {
	if number > 0 {
		fmt.Printf("Number of %s is %.f\n", s, number)
	}
}

func main() {
	var cost, amount_paid float64

	fmt.Print("Enter cost: ")
	_, err := fmt.Scanf("%f", &cost)

	if err != nil {
		panic(err)
	}

	fmt.Print("Enter amount paid: ")
	_, err = fmt.Scanf("%f", &amount_paid)

	if err != nil {
		panic(err)
	}

	if amount_paid < cost {
		fmt.Println("Error! Amount paid is less than cost")
		os.Exit(0)
	} else if amount_paid == cost {
		fmt.Println("Amount paid equals cost...nothing to return")
		os.Exit(0)
	}

	change := amount_paid - cost

	fmt.Println("change = ", change)

	/*calculate bills to return*/
	display(balance(change, 100), "hundred bill(s)")
	change -= balance(change, 100) * 100
	display(balance(change, 50), "fifty bill(s)")
	change -= balance(change, 50) * 50
	display(balance(change, 20), "twenty bill(s)")
	change -= balance(change, 20) * 20
	display(balance(change, 10), "ten bill(s)")
	change -= balance(change, 10) * 10
	display(balance(change, 5), "five bill(s)")
	change -= balance(change, 5) * 5
	display(balance(change, 1), "one bill(s)")
	change -= balance(change, 1) * 1

	//If there's still more change left, return coins
	if change > 0 {
		change *= 100

		fmt.Println("Change = ", change)

		display(balance(change, 25), "quarters")
		change -= balance(change, 25) * 25
		display(balance(change, 10), "dimes")
		change -= balance(change, 10) * 10
		display(balance(change, 5), "nickels")
		change -= balance(change, 5) * 5

		fmt.Printf("Number of pennies is %.f\n", math.Floor(change+0.5))
	}
}
