package main

import "fmt"

func productArray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	result := make([]int, n)

	// Store the product of all elements to the left of each index.
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Multiply by the product of all elements to the right.
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		fmt.Println("i:", i)
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productArray(nums)
	fmt.Println(result) // Output: [24, 12, 8, 6]
}
