/*Reverse a String:
* Enter a string and the program will reverse it and print it out.
*/

package main

import (
		"fmt"
		"os"
		"bufio"
)

func reverse(s string) string{
	str := []rune(s)

	for i := 0; i < len(str)/2; i++ {
		j := len(str) - i - 1
		str[i], str[j] = str[j], str[i]
	}

	return string(str)
}

func main(){
	fmt.Println("Enter string")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()

	fmt.Println(reverse(str))
}
