package main

import "fmt"

func main() {
	arr := []int{5, 6, 8, 4, 2, 9, 8, 2, 4, 1}
	fmt.Println("Initial array is:", arr) //Initial array is: [5 6 8 4 2 9 8 2 4 1]
	fmt.Println("")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted array is: ", arr) //Sorted array is:  [1 2 2 4 4 5 6 8 8 9]
}
