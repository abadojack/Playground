/*Prime Factorization:
* Have the user enter a number and find all Prime Factors (if there are any) and display them.
*/
import java.util.Scanner;

public class Prime {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		int n;
		
		System.out.println("Enter number: ");
		n = input.nextInt();
		
		System.out.println("Prime factors of " + n + " are");
		primeFactors(n);

	}
	
	//primeFactors : finds prime factors of a number num
	public static void primeFactors(int num){
		int tmp = num;
		
		//if 2 is a prime factor of num
		while(tmp%2 == 0){
			System.out.print(2 + " ");
			tmp /= 2;
		}
		
		//find prime factors of n between 3 and sqrt(num)
		for(int i = 3; i <= Math.sqrt(num); i += 2){
			while(tmp%i == 0){
				System.out.print(i + " ");
				tmp /= i;
			}
		}
		
		//when num is a prime number
		if(num == tmp){
			System.out.println(num);
		}
		System.out.println();
	}
	
}
