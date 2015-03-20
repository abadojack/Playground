/*Fizz Buzz:
* Write a program that prints the numbers from 1 to 100. 
* But for multiples of three print “Fizz” instead of the number and for the 
* multiples of five print “Buzz”. For numbers which are multiples of both 
* three and five print “FizzBuzz”.
*/

package main

import (
		"fmt"
)


func main(){
	limit := 100
	num := 1

	for num <= limit {
		if num % 5 == 0 && num % 3 == 0{
			fmt.Println("FizzBuzz")
		}else if num % 3 == 0 {
			fmt.Println("Fizz")
		}else if num % 5 == 0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(num)
		}

		num++
	}
}
