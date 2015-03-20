/* Next Prime Number :
 * Have the program find prime numbers until the user 
 * chooses to stop asking for the next one.
 */
import java.util.Scanner;

public class NextPrimeNumber {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		int choice = 0,i = 1;
		
		//finally found a use for an infinite loop
		while(i++ > 0){
			if(isprime(i)){
				System.out.println("\nPrint next prime number? (0 => no  1 => yes)");
				choice = Integer.parseInt(input.nextLine());
				if(choice == 1){
					System.out.println("\nNext prime number is: " + i);
				}
				else if(choice == 0){
					break;
				}
				else{
					System.out.println("Invalid input...printing next prime number number anyway");
					System.out.println("\nNext prime number is: " + i);
				}
			}
		}
		
	}

	//test if a given number is prime
	private static boolean isprime(int num){
		for(int i = 2; i <= Math.sqrt(num); i++){
			if( num%i == 0 ){
				return false;
			}
		}
		return num > 1;
	}
}
