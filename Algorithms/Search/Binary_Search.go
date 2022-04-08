package main

import "fmt"

func main() {
	numberArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 22, 54, 335}
	x := BinarySearch(52, numberArray)
	fmt.Println(x)
}

// BinarySearch has a O(log^2 N) performance
// This way only returns the value, Returning
// the position in the array and the value is
// implemented in the recursive version.
func BinarySearch(target int, Numbers []int) int {
	found := false
	var result int
	for found != true {
		if len(Numbers) == 0 {
			return -1 // returning -1 for not found
		}
		mid := (len(Numbers) - 1) / 2
		if Numbers[mid] == target {
			found = true
			result = Numbers[mid]
		} else if mid < target {
			Numbers = Numbers[mid+1:]
		} else if mid > target {
			Numbers = Numbers[:mid]
		}

	}
	return result
}
