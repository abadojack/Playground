/*
* Find Cost of Tile to Cover W x H Floor:
* Calculates the total cost of tile it would take to cover a floor plan
* of width and height, using a cost entered by the user.
 */

package main

import "fmt"

func main() {
	var length, width, size, cost, total_cost float64

	fmt.Printf("Length of floor: ")
	fmt.Scanf("%f", &length)

	fmt.Printf("Width of floor: ")
	fmt.Scanf("%f", &width)

	fmt.Printf("Surface area of each tile: ")
	fmt.Scanf("%f", &size)

	fmt.Printf("Cost of each tile: ")
	fmt.Scanf("%f", &cost)

	total_cost = ((length * width) / size) * cost

	fmt.Printf("The total cost is : %.2f\n", total_cost)
}
