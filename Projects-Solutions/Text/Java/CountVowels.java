/*Count Vowels : 
* Enter a string and the program counts the number of vowels in the text.
*  For added complexity have it report a sum of each vowel found.
*/

import java.util.Scanner;

public class CountVowels {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		char[] vowel = {'a', 'e', 'i', 'o', 'u'};	//vowels
		int[] countVowels = {0,0,0,0,0};	//number of each vowels
		int totalCount = 0;	//total number of vowels

		System.out.println("Enter string");
		String str = input.nextLine();
		
		for(int index = 0; index < str.length(); index++){
			int i = 0;
			while(i < vowel.length){
				if(str.charAt(index) == vowel[i]){
					++countVowels[i];
					totalCount++;
				}
				i++;
			}
		}
		
		//display
		int i = 0;
		while( i < vowel.length){
			System.out.println("Sum of " + vowel[i] + " is " + countVowels[i]);
			i++;
		}

		System.out.println("The total number of vowels in the string is " + totalCount);

	}
}

