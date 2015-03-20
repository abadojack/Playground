/* Caesar cipher :
 * Implement a Caesar cipher, both encoding and decoding. The key is an integer
 * from 1 to 25. This cipher rotates the letters of the alphabet (A to Z).
 * The encoding replaces each letter with the 1st to 25th next letter in
 * the alphabet (wrapping Z to A). So key 2 encrypts "HI" to "JK", but key
 * 20 encrypts "HI" to "BC". This simple "monoalphabetic substitution cipher"
 * provides almost no security, because an attacker who has the encoded message
 * can either use frequency analysis to guess the key, or just try all 25 keys.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/* Logic : D(x) = (x - n) mod 26
 * where x is the alphabet being decoded and n is the key.
 * Letters are set to range from 0 t0 26 before the formula is applied i.e.
 * 'A' = 0, 'B' = 1 ... 'Z' = 25 . However, the formula doesn't work in
 * case of a wrap around*/

/*returns decoded string*/
func decode(str string, key int) string {
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if unicode.IsLower(s[i]) {
			s[i] -= 'a'
			s[i] = rune((int(s[i]) - key) % 26)
			if s[i] < 0 {
				s[i] += 'z' + 1
			} else {
				s[i] += 'a'
			}
		} else if unicode.IsUpper(s[i]) {
			s[i] -= 'A'
			s[i] = rune((int(s[i]) - key) % 26)
			if s[i] < 0 {
				s[i] += rune(int('Z') + 1)
			} else {
				s[i] += 'A'
			}
		}
	}
	return string(s)
}

/* Logic : E(x) = (x + n) mod 26
 * where x is the alphabet being encoded and n is the key.
 * Letters are set to range from 0 t0 26 before the formula is applied i.e.
 * 'A' = 0, 'B' = 1 ... 'Z' = 25 */

/* returns encoded string */
func encode(str string, key int) string {
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if unicode.IsLower(s[i]) {
			s[i] -= 'a'
			s[i] = rune((int(s[i]) + key) % 26)
			s[i] += 'a'
		} else if unicode.IsUpper(s[i]) {
			s[i] -= 'A'
			s[i] = rune((int(s[i]) + key) % 26)
			s[i] += 'A'
		}
	}
	return string(s)
}

/*displays all the possible decodings of the string */
func bruteforce(str string) {
	for key := 0; key < 26; key++ {
		fmt.Println("Key >> ", key, " : ", decode(str, key))
	}
}
func main() {
	var choice, key int
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("1.Encode\n2.Decode\n3.Bruteforce")
	_, err := fmt.Scanf("%d", &choice)

	if err != nil {
		panic(err)
	}

	switch choice {
	case 1:
		fmt.Print("Enter key: ")
		_, err = fmt.Scanf("%d", &key)

		if err != nil {
			panic(err)
		}

		if key < 1 || key > 25 {
			fmt.Println("1 <= key <= 25\n")
			os.Exit(1)
		}

		fmt.Println("Enter message to encode")
		input.Scan()
		msg := input.Text()

		fmt.Println("Encoded message is:")
		fmt.Println(encode(msg, key))

		break

	case 2:
		fmt.Print("Enter key: ")
		_, err = fmt.Scanf("%d", &key)

		if err != nil {
			panic(err)
		}

		if key < 1 || key > 25 {
			fmt.Println("1 <= key <= 25")
			os.Exit(1)
		}

		fmt.Println("Enter message to decode")
		input.Scan()
		msg := input.Text()

		fmt.Println("Decoded message: ")
		fmt.Println(decode(msg, key))

		break
	case 3:
		fmt.Println("Enter text to decode")
		input.Scan()
		msg := input.Text()
		bruteforce(msg)
		break

	default:
		fmt.Println("Invalid choice")
	}

}
