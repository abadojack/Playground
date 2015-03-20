/*Sieve of Eratosthenes
* The sieve of Eratosthenes is one of the most
* efficient ways to find all of the smaller primes (below 10 million or so).
 * */

package main

import (
	"fmt"
	"math"
)

func sieve(A []bool, n int) {
	/*A is an slice of boolean values, indexed by integers from 2 to n,
	 * all set to be 'true'*/
	//TODO
	//find a way to set all values in A to true without using a loop
	for i := 2; i < n; i++ {
		A[i] = true
	}

	/*2 is the first prime number, enumerate it's multiples by counting
	 * to n and mark them in the list (i.e. 2i, 3i, e.t.c.) then proceed
	 * to the next unmarked number and do the same*/
	sqrt_n := math.Sqrt(float64(n)) //all factors of an integer are less than it's sqrt
	for i := 2; i < int(sqrt_n); i++ {
		if A[i] {
			p := 2
			for j := i * i; j < n; j = p * i {
				A[j] = false
				p++
			}
		}
	}
}

func main() {
	var n int
	fmt.Println("Input an integer (1 >= n <= 10000000)")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	prime := make([]bool, n)

	sieve(prime, n)

	for i := 0; i < n; i++ {
		if prime[i] {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()

}
