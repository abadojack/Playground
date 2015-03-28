package main

import "fmt"

//Fibonacci Sequence : Enter a number and have the program generate the Fibonacci sequence
//to that number or to the Nth number

func main() {
	var num uint64
	fmt.Println("Enter number:")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic(err)
	}

	f := fib()

	for num > 0 {
		fmt.Print(f(), " ")
		num--
	}

	fmt.Println()
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}
