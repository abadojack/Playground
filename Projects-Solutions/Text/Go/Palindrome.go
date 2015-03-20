/* Check if Palindrome
* Checks if the string entered by the user is a palindrome. 
* That is that it reads the same forwards as backwards like “racecar”
*/

package main

import "fmt"

/*isPalindrome: test if string str is palindrome*/
func isPalindrome(str string)bool{
	start := 0
	end := len(str) -1

	for str[start] == str[end] && end >= start{
		start++
		end--
	}

	if end <= start{
		return true
	}else{
		return false
	}
}

func main(){
	var str string

	fmt.Println("Enter a string to check if it is a palindrome")
	fmt.Scanf("%s", &str)
	
	if isPalindrome(str){
		fmt.Println("\"", str,"\"", " is a palindrome")
	}else{
		fmt.Println("\"", str,"\"", " is not a palindrome")
	}
}
