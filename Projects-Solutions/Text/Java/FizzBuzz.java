/* Fizz Buzz : Write a program that prints the numbers from 1 to 100.
 * But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
 * For numbers which are multiples of both three and five print “FizzBuzz”.
 */
public class FizzBuzz {
    public static void main(String[] args) {
        final int LIMIT = 100;

        for( int num = 1; num <= LIMIT; num++ ){
            if( num%3 == 0 && num%5 == 0 ){
                System.out.println("FizzBuzz");
            }else if( num%3 == 0 ){
                System.out.println("Fizz");
            }else if( num%5 == 0 ){
                System.out.println("Buzz");
            }else{
                System.out.println(num);
            }
        }
    }
}
