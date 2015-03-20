/*Collatz Conjecture** -
* Start with a number *n > 1*.
* Find the number of steps it takes to reach one using the following process:
* If *n* is even, divide it by 2. If *n* is odd, multiply it by 3 and add 1.
 */

package main

import "fmt"

func main() {
	var n int
	count := 0

	fmt.Println("Enter an integer: (n > 1)")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = ((n * 3) + 1)
		}
		count++
		fmt.Printf("Step %d : n == %d\n", count, n)
	}

	fmt.Printf("It took %d steps to reach 1\n", count)
}
