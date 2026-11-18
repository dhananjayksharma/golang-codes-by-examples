package main

import "fmt"

func productArray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Create arrays to store products of elements to the left and right of each index
	left := make([]int, n)
	right := make([]int, n)

	// Fill the left array
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	fmt.Println("Left array:", left) // Debugging statement to print the left array
	
	// Fill the right array
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	fmt.Println("Right array:", right) // Debugging statement to print the right array
	
	// Create the result array by multiplying corresponding elements from left and right arrays
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productArray(nums)
	fmt.Println(result) // Output: [24, 12, 8, 6]
}