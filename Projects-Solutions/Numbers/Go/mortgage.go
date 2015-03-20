/* Mortgage Calculator:
*  Calculate the monthly payments of a fixed term mortgage
*  over given Nth terms at a given interest rate.
*  Also figure out how long it will take the user to pay back the loan
 */

package main

import "fmt"

func main() {
	var months, loan_amount, interest_rate, interest float64

	fmt.Print("Amount of loan: ")
	fmt.Scanf("%f", &loan_amount)

	fmt.Printf("Time in months: ")
	fmt.Scanf("%f", &months)

	fmt.Printf("Interest rate: ")
	fmt.Scanf("%f", &interest_rate)

	interest_rate = (interest_rate / 100)

	interest = (loan_amount * months * interest_rate)

	fmt.Printf("\nInterest : %.2f", interest)
	fmt.Printf("\nTotal money owed : %.2f", (interest + loan_amount))
	fmt.Printf("\nMinimum payment per month:  %.2f\n", (interest+loan_amount)/months)
}
