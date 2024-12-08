package main

import (
	"fmt"
)

func BinarySearch(arr []int, target int) int {
	lower, highest := 0, len(arr)-1

	for lower <= highest {
		mid := lower + (highest-lower)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			highest = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}

func BinarySearchMultiple(arr []int, targets []int) map[int]int {
	results := make(map[int]int)
	for _, target := range targets {
		results[target] = BinarySearch(arr, target)

	}
	return results
}

func main() {
	arr := []int{1, 3, 10, 23, 32, 43, 57, 66, 89, 100}
	fmt.Println(len(arr))
	targets := []int{89}
	results := BinarySearchMultiple(arr, targets)
	for target, index := range results {
		if index != -1 {
			fmt.Printf("Element %d found at index %d\n", target, index)
		} else {
			fmt.Printf("Element %d not found in the array\n", target)
		}
	}

}
