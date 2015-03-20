/*Prime Factorization:
* Have the user enter a number and find all Prime Factors
* (if there are any) and display them.
 */
package main

import (
	"fmt"
	"math"
)

//primefactors : calculates and displays the prime factors of an integer
func primefactors(n int) {
	/* print number of twos that divide n*/
	for n%2 == 0 {
		fmt.Printf("%d ", 2)
		n /= 2
	}

	/*n must be odd hence i += 2
	 * all factor of n are less than sqrt of (n)*/
	sqrt_of_n := math.Sqrt(float64(n))
	for i := 3; i <= int(sqrt_of_n); i += 2 {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}

	/* case where n is a prime number >  2*/
	if n > 2 {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}

func main() {
	var n int

	fmt.Println("Enter an integer to display it's prime factors")
	_, err := fmt.Scanf("%d ", &n)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The prime factors of %d are: \n", n)
	primefactors(n)
}

//check out sieve of Eratosthenes (../../Classic-Algorithms/Go/sieve.go)
