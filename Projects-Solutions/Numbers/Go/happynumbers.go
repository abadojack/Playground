/*Happy Numbers :
* A happy number is defined by the following process. Starting with any positive integer,
* replace the number by the sum of the squares of its digits, and repeat the process
* until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
* Those numbers for which this process ends in 1 are happy numbers, while those that do not end in 1 are unhappy numbers.
* Display an example of your output here. Find first 8 happy numbers.
 */

package main

import "fmt"

/*happy_num: test if a given number is 'happy'*/
func happy_num(num int) bool {
	var digit int
	count := 0
	sum := 0

	for sum != 1 {
		/*limits the number of iterations to 32767*/
		if count > 32767 {
			return false
		}
		digit = num % 10
		num /= 10
		sum += digit * digit
		count++
	}
	return true
}

func main() {
	var num int
	count := 0

	fmt.Println("The first 8 happy numbers are: ")
	for num = 0; num < 32676 && count < 8; num++ {
		if happy_num(num) {
			fmt.Print(num, " ")
			count++
		}
	}
	fmt.Println()
}
