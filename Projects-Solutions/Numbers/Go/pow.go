/* Fast Exponentiation:
 * Ask the user to enter 2 integers a and b and output a^b (i.e. pow(a,b)) in O(lg n) time complexity.
 */
package main

import "fmt"

/*pow: compute a**b using binary powering algorithm */
//TODO
//add comments on how this algorithm works
func pow(a, b int) int {
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

//test the power function
func main() {
	var b, e int

	fmt.Println("Enter base")
	_, err := fmt.Scanf("%d", &b)

	if err != nil {
		panic(err)
	}

	fmt.Println("Enter exponential")
	_, err = fmt.Scanf("%d", &e)

	if err != nil {
		panic(err)
	}

	fmt.Println(pow(b, e))
}
