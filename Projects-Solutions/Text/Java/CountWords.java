/*Count Words in a String : 
* Counts the number of individual words in a string. For added complexity read 
* these strings in from a text file and generate a summary.
*/

import java.util.Scanner;

public class CountWords{
	public static int countWords(String str){
		int index = 0, words = 0;

		while(index < str.length()){
			if(((index > 0) && (str.charAt(index) != ' ') && (str.charAt(index-1) == ' ')) || ((str.charAt(0) != ' ') && (index == 0)))
				words++;
			index++;
		}
		return words;
	}

	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		
		System.out.println("Enter string");
		String line = input.nextLine();

		System.out.println(countWords(line));
	}
}

