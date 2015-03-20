/*
* Fibonacci Sequence
* Enter a number and have the program generate the Fibonacci sequence
* to that number or to the Nth number
 */
package main

import "fmt"

func main() {
	var num uint64
	fmt.Println("Enter number:")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic(err)
	}

	for i := 0; i <= int(num); i++ {
		fmt.Printf("%d ", fibonacci(uint64(i)))
	}

	fmt.Println()
}

func fibonacci(num uint64) uint64 {
	if num == 0 || num == 1 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}
