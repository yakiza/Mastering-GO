package main

import "fmt"

func main() {
	numberArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 22, 54, 335}
	value, index := BinarySearchRecursion(5, numberArray, 0, len(numberArray)-1)
	fmt.Printf("Value %v found on index %v", value, index)
}

// BinarySearchRecursion has a O(log^2 N) performance, this
// version returns the value as well as the index of the
// item in the array
func BinarySearchRecursion(target int, Numbers []int, start, end int) (int, int) {
	if end == start {
		return -1, -1 // -1 represents nothing found
	}
	mid := (start + end) / 2
	if Numbers[mid] == target {
		return Numbers[mid], mid
	} else if mid < target {
		return BinarySearchRecursion(target, Numbers, mid+1, end)
	} else {
		return BinarySearchRecursion(target, Numbers, start, mid-1)
	}

}
