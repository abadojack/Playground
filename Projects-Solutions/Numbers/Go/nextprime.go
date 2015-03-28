/*
 * Next Prime Number :
 * Have the program find prime numbers until the user chooses to stop
 * asking for the next one.
 */

package main

import (
	"fmt"
	"math"
)

//Returns true if it's argument is prime or false otherwise
func isPrime(num int) bool {
	sqrt_num := math.Sqrt(float64(num))
	for i := 2; i <= int(sqrt_num); i++ {
		if num%i == 0 {
			return false
		}
	}

	return num > 1
}

func main() {
	choice := "yes"
	for i := 2; i < math.MaxInt64 && choice == "y" || choice == "yes"; i++ {
		if isPrime(i) {
			fmt.Println(i)
			fmt.Println("Print next prime number? (y)es/(n)o")
			_, err := fmt.Scanf("%s", &choice)

			if err != nil {
				panic(err)
			}
		}
	}
}
