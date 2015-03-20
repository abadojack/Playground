/*Collatz Conjecture** 
* Start with a number *n > 1*. Find the number of steps it takes to reach one using the 
* following process: If *n* is even, divide it by 2. If *n* is odd, multiply it by 3 and add 1.
*/
import java.util.Scanner;

public class CollatzConjecture {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		long count = 0;
		
		System.out.println("Enter number: ");
		int n = Integer.parseInt(input.nextLine());
		
		if(n < 0){
			System.out.println("Enter a positive number");
			n = Integer.parseInt(input.nextLine());
		}else{
			while(n != 1){
				if(n%2 == 0){
					n /= 2;
					count++;
				}else{
					n = (n * 3) + 1;
					count++;
				}
			}
		}
		
		System.out.println("It took " + count + " steps to reach 1");
	}
}
