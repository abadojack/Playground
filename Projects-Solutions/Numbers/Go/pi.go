/*
* Find PI to the Nth Digit :
* Enter a number and have the program generate PI up to that many decimal places.
* Keep a limit to how far the program will go
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Println("To how many decimal places do you want to compute PI?")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	if n >= 0 && n <= 48 {
		fmt.Printf("%.*f\n", n, math.Pi)
	} else {
		fmt.Println("Limits : 0 >= n <= 48")
	}
}
