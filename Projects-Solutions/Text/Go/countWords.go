/*Count Words in a String : 
* Counts the number of individual words in a string. For added complexity read 
* these strings in from a text file and generate a summary.
*/

package main

import (
		"os"
		"bufio"
		"fmt"
)

func countWords(str string)int {
	index := 0
	words := 0

	for index < len(str) {
		if ((index > 0) && (str[index] != ' ') && (str[index-1] == ' ')) || ((str[0] != ' ') && (index == 0) ){
			words++
		}
		index++
	}

	return words
}
func main(){
	fmt.Println("Enter string")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	fmt.Println(countWords(str))
}

