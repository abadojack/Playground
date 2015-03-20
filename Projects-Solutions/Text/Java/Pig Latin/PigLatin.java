/*Class PigLatin : contains functions that deal with converting words to Pig Latin*/

public class PigLatin {

    //returns true if character 'ch' is a vowel
    public static boolean isVowel(char ch){
        char[] vowels = {'a', 'e', 'i', 'o', 'u'};	//vowels
        boolean flag = false;
        int index = 0;
        while(index < vowels.length){
            if(ch == vowels[index]){
                flag = true;
                break;
            }
            index++;
        }
        return flag;
    }


    /*For words that begin with consonant sounds, the initial consonant or
	* consonant cluster is moved to the end of the word, and "ay"
	* "banana" → "ananabay"
	* "glove" → "oveglay"*/
    public static String beginConsonant(String word){
        int index = 0;
        while( !isVowel(word.charAt(index)) ){
            index++;
        }

        return word.substring(index, word.length()) + word.substring(0, index) + "ay";
    }


    /*For words that begin with vowel sounds or silent letter, you just
	 * add "way" (or "wa") to the end.
	 * "egg" → "eggway"
	 *"inbox" → "inboxway"
	 */
    public static String beginVowel(String word){
        return word + "way";
    }


    //converts String to pig Latin
    public static String toPigLatin(String word){
        if(isVowel(word.charAt(0))){
            return beginVowel(word);
        }else{
            return beginConsonant(word);
        }
    }

}
