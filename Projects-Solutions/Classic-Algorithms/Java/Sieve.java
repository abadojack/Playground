/*
* Sieve of Eratosthenes:
* The sieve of Eratosthenes is one of the most efficient ways to
* find all the smaller primes(below 10 million or so).
*/

import java.util.Scanner;

public class Sieve {
	public static void sieve(boolean[] A, int n){
		/*A is indexed by integers from 2 to n,
		* all set to be 'true'*/
		for(int i = 2; i < n; i++){
			A[i] = true;
		}

		/*2 is the first prime number, enumerate it's multiples by counting
		* to n and mark them in the list (i.e. 2i, 3i, e.t.c) then proceed
		* to the next unmarked number and do the same */
		for(int i = 2; i < Math.sqrt(n); i++){
			if(A[i]){
				int p = 2;
				for(int j = i * i; j < n; j = p++ * i){
					A[j] = false;
				}
			}
		}
	}

	//test the method sieve
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		int n = 0;
		boolean[] prime = new boolean[n];
		
			
		try{
			System.out.println("Enter an integer");
			n = Integer.parseInt(input.nextLine());
			prime = new boolean[n];
		}catch(NegativeArraySizeException exception){
			n = Math.abs(n);
			prime = new boolean[n];
		}catch(Exception exception){
			exception.printStackTrace();
		}finally{

			sieve(prime, n);
			
			//print prime numbers between 1 and n
			for(int i = 0; i < n; i++){
				if(prime[i]){
					System.out.print(i + " ");
				}
			}

			System.out.println();
		}
	}
}
