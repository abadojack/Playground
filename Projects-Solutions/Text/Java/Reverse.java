/* Reverse a String
* Enter a string and the program will reverse it and print it out.*/

import java.util.Scanner;

public class  Reverse {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Enter string to reverse");
        String str = input.nextLine();

        StringBuffer strBuf = new StringBuffer(str);

        strBuf.reverse();

        System.out.println(strBuf);
    }
}
