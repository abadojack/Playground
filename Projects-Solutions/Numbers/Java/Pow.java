/* Fast Exponentiation:
 * Ask the user to enter 2 integers a and b and output a^b (i.e. pow(a,b)) in O(lg n) time complexity.
 */

import java.util.Scanner;

public class Pow {
	public static double power(double base, int exponential){
		double temp;
		//base case
		if(exponential == 0){
			return 1;
		}

		temp = power(base, exponential/2);

		//exponential is even
		if(exponential%2 == 0){
			return temp * temp;
		}else{
			//exponential is odd
			if(exponential > 0){
				//exponential is odd and positive
				return base * temp * temp;
			}else{
				//exponential is negative
				return (temp * temp)/base;
			}

		}
	}
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		
		System.out.println("Enter base");
		int b = Integer.parseInt(input.nextLine());

		System.out.println("Enter exponential");
		int exp = Integer.parseInt(input.nextLine());

		System.out.println(power(b, exp));
	}
}
