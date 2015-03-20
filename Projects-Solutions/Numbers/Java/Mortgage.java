/*Mortgage Calculator:
* Calculate the monthly payments of a fixed term mortgage over given Nth terms at a given interest rate. 
* Also figure out how long it will take the user to pay back the loan.
*/

import java.util.Scanner;

public class Mortgage {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		double loan_amount, interest, interest_rate;
		int months;
		
		System.out.println("Amount loaned");
		loan_amount = input.nextDouble();
		
		System.out.println("Time in months");
		months = input.nextInt();
		
		System.out.println("Interest rate");
		interest_rate = input.nextDouble();
		
		interest_rate = (interest_rate/100);
		interest = loan_amount * interest_rate * months;
		
		System.out.println("Interest: " + interest);
		System.out.println("Total money owed: " + (interest + loan_amount));
		System.out.println("Minimum payment per: " + (interest + loan_amount)/months);
	}
}
