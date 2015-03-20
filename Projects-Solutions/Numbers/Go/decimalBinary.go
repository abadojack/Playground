/*Binary to Decimal and Back Converter :
* Develop a converter to convert a decimal
* number to binary or a binary number to its decimal equivalent.
 */

package main

import "fmt"

/*Integer power: compute a**b using binary powering algorithm */
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

/*isBinary : returns true if it's argument is a binary digit else false otherwise*/
func isBinary(n int) bool {
	for n > 0 {
		digit := n % 10
		n /= 10
		flag := digit == 1 || digit == 0
		if !flag {
			return false
		}
	}

	return true
}

/* converts decimal to binary*/
func dec2bin(n int) int {
	var rem, i, bin int
	bin = 0
	i = 1

	for n != 0 {
		rem = n % 2
		n /= 2
		bin += rem * i
		i *= 10
	}
	return bin
}

/* converts binary to decimal*/
func bin2dec(n int) int {
	var rem, dec, i int
	dec = 0
	i = 0

	for n != 0 {
		rem = n % 10
		n /= 10
		dec += rem * Pow(2, i)
		i++
	}
	return dec
}

func main() {
	var num, choice int
	fmt.Print("1.decimal to binary\n2.binary to decimal\n\nEnter choice: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("Enter decimal number")
		fmt.Scanf("%d", &num)
		bin := dec2bin(num)

		fmt.Printf("%d to binary is %d\n", num, bin)

		break

	case 2:
		fmt.Println("Enter binary number")
		fmt.Scanf("%d", &num)

		if isBinary(num) {
			bin := bin2dec(num)
			fmt.Printf("%d to binary is %d\n", num, bin)
		} else {
			fmt.Println(num, " is not a binary number")
		}

		break

	default:
		fmt.Println("Invalid choice")
	}
}
