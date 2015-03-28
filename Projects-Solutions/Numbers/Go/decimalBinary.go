package main

import "fmt"
import "os"
import "strconv"

//Binary to Decimal and Back Converter :
//Develop a converter to convert a decimal number to binary or a binary number to its decimal equivalent.

//Compute a**b using binary powering algorithm
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

//Return true if the argument is a binary digit else false otherwise
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

//Convert decimal number to its binary equivalent
func dec2bin(n int) int {
	bin := 0
	i := 1
	for n != 0 {
		bin += i * (n % 2)
		n /= 2
		i *= 10
	}
	return bin
}

//convert binary number to its decimal equivalent
func bin2dec(n int) int {
	i, dec := 0, 0
	for n != 0 {
		rem := n % 10
		dec += (Pow(2, i) * rem)
		n /= 10
		i++
	}
	return dec
}

func usage() {
	fmt.Println("usage: ")
	fmt.Println(os.Args[0], " -option <integer>\n")
	fmt.Println("options : -b to convert  to binary or -d to convert to decimal\n")
	fmt.Println(os.Args[0], " -b 2\n")
	fmt.Println(os.Args[0], "-d 1010\n")
	os.Exit(0)
}

func main() {
	var num int

	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "-d":

		if isBinary(num) {
			fmt.Println(bin2dec(num))
		} else {
			usage()
		}
		break

	case "-b":
		fmt.Println(dec2bin(num))
		break

	default:
		usage()
	}
}
