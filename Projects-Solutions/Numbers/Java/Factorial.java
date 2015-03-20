/*Factorial Finder :
* The Factorial of a positive integer, n, is defined as the product of the sequence n, n-1, n-2, ...1 
* and the factorial of zero, 0, is defined as being 1. Solve this using both loops and recursion.
*/

import java.math.BigInteger;
import java.util.Scanner;

public class Factorial {
	//factorial by recursion
	public static BigInteger factorial(int n){
		if(BigInteger.valueOf(n).equals(BigInteger.ONE) || BigInteger.valueOf(n).equals(BigInteger.ZERO)){
			return BigInteger.ONE;
		}else{
			return BigInteger.valueOf(n).multiply(factorial(n - 1));
		}
	}
	
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		BigInteger fact = BigInteger.ONE;
		
		System.out.println("Enter an integer");
		int num = input.nextInt();
		
		if(num < 0){
			System.out.println("Input should be a positive integer\nExiting...");
			System.exit(1);
		}
		
		//factorial using a loop
		for(int i = num; i > 0; i--){
			fact = fact.multiply(BigInteger.valueOf(i));
		}
		
		System.out.println("\nUsing a loop, the factorial is: ");
		System.out.println(fact);
		
		System.out.println("\nBy recursion, the factorial is: ");
		System.out.println(factorial(num));
	}
}
