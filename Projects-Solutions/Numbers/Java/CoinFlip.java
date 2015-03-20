/*Coin Flip Simulation** -
*Write some code that simulates flipping a single coin however many times the user decides.
*The code should record the outcomes and count the number of tails and heads.
*/

import java.util.Random;
import java.util.Scanner;

public class CoinFlip {
	private static Random randomInt = new Random(System.currentTimeMillis());
	
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		int heads = 0, tails = 0;
		
		System.out.println("How many times do you want to flip the coin?");
		int num = input.nextInt();
		
		for(int i = 0; i < num; i++){
			if(flip() == 1){
				heads++;
			}else{
				tails++;
			}
		}
		
		System.out.println("You flipped heads " + heads + " times");
		System.out.println("You flipped tails " + tails + " times");
	}
	
	/*returns 1 or 2 randomly*/
	public static int flip(){
		return 1 + randomInt.nextInt(2);
	}
}
