/*Find e to the Nth Digit - 
Just like the previous problem, but with e instead of PI. Enter a number and have the
program generate e up to that many decimal places. Keep a limit to how far the program will go.

Solution logic:
2.71828 * Math.pow(10,3) = 2718.28
Math.floor(2718.28) = 2718
2718/Math.pow(10,3) = 2.718
*/
import java.util.Scanner;

public class E {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		final double E = Math.E;
		
		System.out.println("To what decimal place do you want to view E?");
		int n = input.nextInt();
		
		if(n > 15){
			System.out.println("Limit Exceeded\n");
		}else{
			System.out.println(Math.floor(E * Math.pow(10, n))/Math.pow(10, n));
		}
	}
}
