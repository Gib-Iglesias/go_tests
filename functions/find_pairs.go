package main

import (
	"fmt"
)

func findPairsWithSum(arr1, arr2 []int, x int) {
	arr2Map := make(map[int]bool)
	for _, num := range arr2 {
		arr2Map[num] = true
	}

	fmt.Printf("Pairs for X = %d:\n", x)
	for _, num1 := range arr1 {
		complement := x - num1
		if arr2Map[complement] {
			fmt.Printf("%d %d\n", num1, complement)
		}
	}
}

func main() {
	// Example 1
	arr1 := []int{-1, -2, 4, -6, 5, 7}
	arr2 := []int{6, 3, 4, 0}
	x1 := 8
	findPairsWithSum(arr1, arr2, x1)
	fmt.Println("--------------------")
	// Example 2
	arr3 := []int{1, 2, 4, 5, 7}
	arr4 := []int{5, 6, 3, 4, 8}
	x2 := 9
	findPairsWithSum(arr3, arr4, x2)
	fmt.Println("--------------------")
}
