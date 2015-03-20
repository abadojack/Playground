/*Factorial Finder** - The Factorial of a positive integer, n,
* is defined as the product of the sequence n, n-1, n-2, ...1 and the factorial of zero, 0,
* is defined as being 1. Solve this using both loops and recursion.
 */

package main

import "fmt"

/* factorial by recursion*/
func fact_recursion(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact_recursion(n-1)
	}
}

/* factorial by loop*/
func fact_loop(n uint) uint {
	var factorial uint = 1

	for ; n > 0; n-- {
		factorial *= n
	}

	return factorial
}

func main() {
	var num uint
	fmt.Print("Enter number find it's factorial: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic(err)
	}

	if num >= 0 && num <= 65 {
		fmt.Printf("Factorial of %d by recursion is %d\n", num, fact_recursion(num))
		fmt.Printf("Factorial of %d by loops is %d\n", num, fact_loop(num))
	} else {
		fmt.Println("Limits : 0 >= num <= 65")
	}
}
