/* Sorting: Implement two types of sorting algorithms:
*  Merge sort and bubble sort.
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int

	fmt.Print("Enter length of array to be sorted (<=1024) : ")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	if n > 1024 || n < 1 {
		fmt.Println("Limits : 1 < n < 1024")
	}

	fmt.Println("\n\t\tBUBBLE SORT")
	arr1 := createArray(n)
	fmt.Println("\n\tArray before bubble sort:")
	fmt.Println(arr1)
	fmt.Println("\n\tBubble sorted array:")
	fmt.Println(bubblesort(arr1))

	fmt.Println("\n\t\tMERGE SORT")
	arr2 := createArray(n)
	fmt.Println("\n\tArray before merge sort:")
	fmt.Println(arr2)
	mergeSort(arr2, 0, n-1)
	fmt.Println("\n\tMerge sorted array: ")
	fmt.Println(arr2)
}

/*bubblesort: sort array by comparison and exchange*/
func bubblesort(n []int) []int {
	sorted := false

	for i := 0; i < len(n) && !sorted; i++ {
		sorted = true
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[j+1] {
				temp := n[j]
				n[j] = n[j+1]
				n[j+1] = temp
				sorted = false
			}
		}
	}
	return n
}

/* Function to merge the two haves arr[l..m] and arr[m+1..r] of array arr[] */
func merge(arr []int, l, m, r int) {
	var i, j, k int
	n1 := m - l + 1
	n2 := r - m

	/* create temp arrays */
	R := make([]int, n2)
	L := make([]int, n1)

	/* Copy data to temp arrays L[] and R[] */
	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	/* Merge the temp arrays back into arr[l..r]*/
	i = 0
	j = 0
	k = l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	/* Copy the remaining elements of L[], if there are any */
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy the remaining elements of R[], if there are any */
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

/* l is for left index and r is right index of the sub-array
of arr to be sorted */
func mergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2 //Same as (l+r)/2, but avoids overflow for large l and h
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

/*function to create an array of length n from randomly generated integers
 * between 0 - 255*/
func createArray(n int) []int {

	rand.Seed(time.Now().UTC().UnixNano())

	arr := make([]int, n)

	for i := range arr {
		arr[i] = rand.Intn(255)
	}

	return arr
}
