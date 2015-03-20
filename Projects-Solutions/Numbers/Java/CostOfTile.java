/*
* Find Cost of Tile to Cover W x H Floor:
* Calculates the total cost of tile it would take to cover a floor plan
* of width and height, using a cost entered by the user.
 */

import java.util.Scanner;

public class CostOfTile {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);

		System.out.println("Enter length of the floor: ");
		double len = input.nextDouble();

		System.out.println("Enter width of the floor: ");
		double width = input.nextDouble();

		System.out.println("Surface area of each tile: ");
		double tileSize = input.nextDouble();

		System.out.println("cost of each tile: ");
		double cost = input.nextDouble();

		double total_cost = ((len * width) / tileSize) * cost;

		System.out.println("The total cost is : " + total_cost);
	}

}
