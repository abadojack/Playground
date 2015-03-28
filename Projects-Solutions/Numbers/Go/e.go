package main

import (
	"fmt"
	"math"
)

// Find e to the Nth Digit :Just like the previous problem, but with e instead of PI.
//Enter a number and have the program generate e up to that many decimal places.Keep a limit to how far the program will go.

func main() {
	var n int

	fmt.Println("To how many terms do you want to generate e ?")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	if n >= 0 && n <= 51 {
		fmt.Printf("%.*f\n", n, math.E)
	} else {
		fmt.Println("Limits: 0 >= n <= 51")
	}
}
