package main

import "fmt"

func findDuplicate(nodes []string) []string {
	// Create a map to store the count of each string
	countMap := make(map[string]int)
	for _, node := range nodes {
		countMap[node]++
	}

	finalUnique := make([]string, 0)
	for node, count := range countMap {
		if count > 1 {
			finalUnique = append(finalUnique, node)
		}
	}
	return finalUnique
}
func main() {
	// Example usage of the findDuplicate function
	nodes := []string{"a", "b", "a", "d", "e", "c", "a", "d", "b", "j"}
	duplicate := findDuplicate(nodes)
	fmt.Println("The duplicate strings are:", duplicate)
}
