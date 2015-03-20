/*Fibonacci Sequence
* Enter a number and have the program generate the Fibonacci sequence
* to that number or to the Nth number.
*/

import java.math.BigInteger;
import java.util.Scanner;

public class Fibonacci {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		System.out.println("To what number do you want to generate the fibonacci series?");
		double num = input.nextInt();

		for(int i = 0; i <= num; i++){
			System.out.println(fibonacci(BigInteger.valueOf(i)));
		}
	}

	public static BigInteger fibonacci(BigInteger number){
		BigInteger TWO = BigInteger.valueOf(2);
		if(number.equals(BigInteger.ONE) || number.equals(BigInteger.ZERO)){	//base case
			return number;
		}else{
			return fibonacci(number.subtract(BigInteger.ONE)).add(fibonacci(number.subtract(TWO)));
		}
	}
}
