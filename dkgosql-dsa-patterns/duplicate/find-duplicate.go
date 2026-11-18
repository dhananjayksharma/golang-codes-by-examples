package main

import "fmt"

func findDuplicates(arr []string) []string {
	seen := make(map[string]int, len(arr))
	final := make([]string, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if j, exists := seen[arr[i]]; exists {
			seen[arr[i]] = j + 1
		} else {
			seen[arr[i]] = 1
		}
		if seen[arr[i]] > 1 {
			final = append(final, arr[i])
		}
	}
	return final
}

func main() {
	arr := []string{"b", "c", "c", "d", "b", "a", "d", "a", "a", "b"}
	duplicates := findDuplicates(arr)
	fmt.Println("Duplicates:", duplicates)
}
