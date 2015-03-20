import java.util.Random;

public class sort {

    public static void main(String[] args) {
        Random randInt = new Random(System.nanoTime());
        Sort sortArr = new Sort();
        int[] arr = new int[20];
        int[] arrCopy2 = new int[arr.length];
        int[] arrCopy1 = new int[arr.length];

        /*create an array of random integers*/
        for(int i = 0; i < arr.length; i++) {
            arr[i] = 1 + randInt.nextInt(100);
            arrCopy1[i] = arr[i];
            arrCopy2[i] = arr[i];
        }

        System.out.println("Bubble sort\nBefore bubble sort");
        displayArray(arr);
        sortArr.bubbleSort(arr);
        System.out.println("After bubble sort");
        displayArray(arr);

        System.out.println("\n\nMerge sort\nBefore merge sort");
        displayArray(arrCopy1);
        sortArr.mergeSort(arrCopy1, 0, arrCopy1.length - 1);
        System.out.println("After merge sort");
        displayArray(arrCopy1);

        System.out.println("\n\nQuick sort\nBefore quick sort sort");
        displayArray(arrCopy2);
        sortArr.quickSort(arrCopy2, 0, arrCopy2.length - 1);
        System.out.println("After quick sort");
        displayArray(arrCopy2);
    }

    public static void displayArray(int arr[]){
        for(int x: arr){
            System.out.print(x + " ");
        }
        System.out.println(" ");
    }
}

class Sort {
    //quick sort
    public void quickSort(int[] list, int start, int end){
        int i = start;
        int j = end;
        int pivot = list[(start + end)/2];

        if( i <= j ){
            while( list[i] < pivot ){
                i++;
            }

            while( list[j] > pivot ){
                j--;
            }
        }

        if(i <= j ){
            int tmp = list[i];
            list[i] = list[j];
            list[j] = tmp;
            i++;
            j--;
        }

        if(start < j ){
            quickSort(list, start, j);
        }

        if(end > i){
            quickSort(list, i, end);
        }

    }


    //merge sort
    public void mergeSort( int array[], int start, int end) {
        int middle = 0, left, right, temp;

        if( start < end ){
            //split the array in half and sort each half
            middle = (start + end)/2;
            mergeSort(array, start, middle);
            mergeSort(array, middle + 1, end);
        }

        //merge the sorted arrays into one
        left = start;
        right = middle + 1;

        //while there are integer in the array to be sorted
        while(left <= middle && right <= end){
            /*if the current integer in the left array is larger the current integer in the right array
            * the integers need to be swapped*/
            if(array[left] > array[right]){
                temp = array[right];

                /*move the left array right one position to make room for the smaller integer*/
                for(int i = right - 1; i >= left; i--){
                    System.arraycopy( array, i, array, i + 1, 1 );
                }

                //put the smaller integer where it belongs
                array[left] = temp;

                //the right array and the middle need to shift right
                right++;
                middle++;
            }

            //no matter what the left array moves right
            left++;
        }
    }

    //bubble sort
    public void bubbleSort(int list[]){
        for(int i = 0; i <= list.length - 1; i++ ){
            for(int j = 0; j < list.length - 1; j++  ){
                if(list[j] > list[j+1]){
                    int temp = list[j];
                    list[j] = list[j + 1];
                    list[j+1] = temp;
                }
            }
        }
    }
}
