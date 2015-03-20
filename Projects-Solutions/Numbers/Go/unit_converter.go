/* Unit Converter (temp, volume, mass and more)
* Converts various units between one another.
* The user enters the type of unit being entered,
* the type of unit they want to convert to and then the value.
* The program will then make the conversion.
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	clearscreen()
	choose()
	continue_or_exit()
}

/*temperature: temperature conversion*/
func temperature() {
	var choice int
	var temp float64
	fmt.Println("1. Celsius to fahrenheit\n2. Fahrenheit to celsius")
	fmt.Print("? ")
	choice = getInt()

	switch choice {
	/* celsius to fahrenheit*/
	case 1:
		temp = getInput()
		temp = 9.0/5.0*temp + 32
		break

	/* fahrenheit to celsius */
	case 2:
		temp = getInput()
		temp = 5.0 / 9.0 * (temp - 32)
		break

	default:
		fmt.Println("Invalid choice")
		break
	}
	fmt.Println(temp)
}

/* mass : mass/weight convertion */
func mass() {
	var choice int
	var weight float64
	fmt.Println("1. Ounces to grams\n2. Grams to ounces\n3. Kilogram to pounds\n4. Pounds to kilograms")
	fmt.Printf("? ")
	choice = getInt()

	switch choice {
	/* ounces to grams */
	case 1:
		weight = convert(28.35)
		break
	/* grams to ounces*/
	case 2:
		weight = convert(0.03527)
		break
		/* kilograms to pounds*/
	case 3:
		weight = convert(2.2045)
		break
	/* pounds to kilograms*/
	case 4:
		weight = convert(0.4536)
		break

	default:
		fmt.Println("Invalid choice")
		break
	}
	fmt.Println(weight)
}

/*volume : volume conversion*/
func volume() {
	var choice int
	var vol float64
	fmt.Println("1.Inches^3 to cm^3\n2.cm^3 to inches^3\n3.Feet^3 to metres^3\n4.metres^3 to feet^3")
	choice = getInt()

	switch choice {
	/* inches3 to cm3*/
	case 1:
		vol = convert(16.38706)
		break

	/*cm3 to inches3*/
	case 2:
		vol = convert(0.061024)
		break

	/* feet3 to m3*/
	case 3:
		vol = convert(0.02832)
		break

	/* m3 to feet3*/
	case 4:
		vol = convert(35.31467)
		break

	default:
		fmt.Println("invalid choice")
		break
	}
	fmt.Println(vol)
}

/* length: linear conversion */
func convert_length() {
	var choice int
	var length float64
	fmt.Println("1.Inches to cm\n2.cm to Inches\n3.Feet to metres\n4.Metres to feet")
	fmt.Println("5.Yards to metres\n6.Metres to yards\n7.Miles to km\n8.Km to miles")
	choice = getInt()

	switch choice {
	/*inches to cm*/
	case 1:
		length = convert(2.54)
		break
		/*cm to inches*/
	case 2:
		length = convert(0.3937)
		break
	/*feet to metres*/
	case 3:
		length = convert(0.3048)
		break
	/*metres to feet*/
	case 4:
		length = convert(3.28084)
		break
	/*yards to metres*/
	case 5:
		length = convert(0.9144)
		break
	/*metres to yards*/
	case 6:
		length = convert(1.09613)
		break
	/*miles to km*/
	case 7:
		length = convert(1.60934)
		break
	/*km to miles*/
	case 8:
		length = convert(0.621371)
		break
	default:
		fmt.Println("Invalid choice")
		break

	}
	fmt.Println(length)
}

/* choose : choose conversion to perform*/
func choose() {
	var choice int
	clearscreen()
	fmt.Println("1.Convert length\n2.Convert mass\n3.Convert Temperature\n4.Convert Volume\n5.exit")
	fmt.Println("? ")
	choice = getInt()

	switch choice {
	case 1:
		convert_length()
		break
	case 2:
		mass()
		break
	case 3:
		temperature()
		break
	case 4:
		volume()
		break
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
		break
	}

}

/* prompts user to exit or continue running the program*/
func continue_or_exit() {
	var choice int
	fmt.Println("\n\n1.continue\n2.exit")
	choice = getInt()
	switch choice {
	case 1:
		choose()
		break
	case 2:
		os.Exit(0)
		break
	default:
		fmt.Println("Invalid choice")
		break

	}
	clearscreen()
}

/* clears screen using ANSI escape sequences*/
func clearscreen() {
	fmt.Printf("\033[2J")
	fmt.Printf("\033[%d;%dH", 0, 0)
}

/* getInt : read integer input from user */
func getInt() int {
	var num int
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		panic(err)
	}

	return num
}

/* read value to convert from user */
func getInput() float64 {
	var num float64

	fmt.Print("Enter value to convert : ")
	_, err := fmt.Scanf("%f", &num)

	if err != nil {
		panic(err)
	}

	return num
}

/*convert: read value to convert from user and multiply it with the 'conversion constant' */
func convert(constant float64) float64 {
	var num float64

	fmt.Print("Enter value to convert : ")
	_, err := fmt.Scanf("%f", &num)

	if err != nil {
		panic(err)
	}

	num *= constant

	return num
}
