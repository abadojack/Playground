/*Pig Latin** - Pig Latin is a game of alterations played on the English
* language game. To create the Pig Latin form of an English word the initial
* consonant sound is transposed to the end of the word and an ay is affixed 
* (Ex.: "banana" would yield anana-bay). Read Wikipedia for more information
* on rules.*/

package main

import (
		"fmt"
		"strings"
		"os"
		"bufio"
)


/*returns true if 'ch' is a vowel*/
func isVowel(ch byte)bool{
	//convert 'ch' to lowercase
	 tmp := []byte(strings.ToLower(string(ch)))
	 ch = tmp[0]
	
	vowels := []byte("aeiou")

	for _, x := range vowels {
		if x == ch {
			return true
		}
	}
	return false
}

/*extract bytes from 'source' beginning from index 'sourceStart' to 'sourceEnd'
 * into 'target' and return 'target' */
func getBytes(source []byte, sourceStart, sourceEnd int)[]byte{
	target := make([]byte, sourceEnd - sourceStart)
	i := 0
	for index := sourceStart; index < sourceEnd; index++ {
		target[i] = source[index]
		i++
	}
	return target
}

/*For words that begin with consonant sounds, the initial consonant or
* consonant cluster is moved to the end of the word, and "ay"
* "banana" → "ananabay"
* "glove" → "oveglay"*/
func wordsThatBeginWithConsonant(word string)string {
	index := 0
	wordInBytes := []byte(word)
	
	//find number of consonants that prefix the word
	for !isVowel(wordInBytes[index]){
		index++
	}
	
	
	return (string(getBytes(wordInBytes,index, len(wordInBytes))) + string(getBytes(wordInBytes, 0, index)) + "ay")
}

/*For words that begin with vowel sounds or silent letter, you just
* add "way" (or "wa") to the end.
* "egg" → "eggway"
* "inbox" → "inboxway" */
func wordsThatBeginWithVowel(word string)string{
	return word + "way"
}

/*returns 'word' in Pig Latin*/
func toPigLatin(word string) string{
	wordInBytes := []byte(word)
	if isVowel(wordInBytes[0]){
		return wordsThatBeginWithVowel(word)
	}else{
		return wordsThatBeginWithConsonant(word)
	}
}


func main(){
	fmt.Println("Enter word to convert to pig Latin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	
	fmt.Println(toPigLatin(input.Text()))
}
