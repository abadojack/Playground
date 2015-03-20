/*Count Vowels: 
* Enter a string and the program counts the number of vowels in the text. 
* For added complexity have it report a sum of each vowel found.*/

package main

import (
		"fmt"
		"strings"
		"os"
		"bufio"
)

func main() {
	vowels := []string{"a", "e", "i", "o", "u"}
	noOfVowels := 0

	fmt.Println("Enter text")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.ToLower(input.Text())

	for _, v := range vowels{
		noOfVowels += strings.Count(str, v)
		fmt.Println(v, " : ", strings.Count(str, v))
	}

	fmt.Println("Number of vowels in text is ", noOfVowels)
}
