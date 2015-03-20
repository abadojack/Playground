/*Calculator:
* A simple calculator to do basic operators. Make it a scientific
* calculator for added complexity.
* */

package main

import (
	"fmt"
	"math"
)

//performs basic arithmetic ... addition, subtraction, division and multiplication
func arithmetic() float64 {
	var num1, num2, ans float64
	operator := '+'
	fmt.Println("Enter input as num1(operator)num2 e.g. 6+7: ")
	fmt.Scanf("%f %c %f", &num1, &operator, &num2)

	switch operator {
	case '+':
		ans = num1 + num2
		break

	case '-':
		ans = num1 - num2
		break

	case '*':
		ans = num1 * num2
		break

	case '/':
		ans = num1 / num2
		break
	default:
		fmt.Println("Invalid output")
		break
	}

	return ans
}

//performs sin,cos,tan,log,sqrt
func scientific_calc() float64 {
	var num, answer float64
	var choice int

	fmt.Print("1.sin\n2.cos\n3.tan\n4.log\n5.square root\nchoice:\n?")
	_, err := fmt.Scanf("%d", &choice)

	if err != nil {
		fmt.Println("Error:", err)
	}

	switch choice {
	case 1:
		fmt.Println("Enter the number: ")
		fmt.Scanf("%f", &num)
		answer = math.Sin(num)
		break
	case 2:
		fmt.Println("Enter the number: ")
		fmt.Scanf("%f", &num)
		answer = math.Cos(num)
		break
	case 3:
		fmt.Println("Enter the number: ")
		fmt.Scanf("%f", &num)
		answer = math.Tan(num)
		break
	case 4:
		fmt.Println("Enter the number: ")
		fmt.Scanf("%f", &num)
		answer = math.Log(num)
		break
	case 5:
		fmt.Println("Enter the number: ")
		fmt.Scanf("%f", &num)
		answer = math.Sqrt(num)
	default:
		fmt.Println("Invalid input")
		break
	}
	return answer
}

func main() {
	var ans float64
	var choice int

	fmt.Println("1. Basic Calcualtor")
	fmt.Println("2. Scientic Calculator")
	_, err := fmt.Scanf("%d", &choice)

	if err != nil {
		fmt.Println("Error", err)
	}

	switch choice {
	case 1:
		ans = arithmetic()
		break
	case 2:
		ans = scientific_calc()
		break
	default:
		fmt.Println("Invalid choice")
		break
	}

	fmt.Println(ans)
}
