/*Check if Palindrome:
* Checks if the string entered by the user is a palindrome. That is that it reads the same forwards as backwards like “racecar”
*/

import java.util.Scanner;

public class Palindrome{
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);

		System.out.println("Enter string");
		String str = input.nextLine();

		StringBuffer strCopy = new StringBuffer(str);

		strCopy.reverse();

		String reversedString = new String(strCopy);
		

		if(str.equals(reversedString)){
			System.out.println("'" + str + "'" + " is a palindrome");
		}else{
			System.out.println("Not a palindrome");
		}
	}
}

