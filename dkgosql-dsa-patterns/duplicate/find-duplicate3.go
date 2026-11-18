package main

import "fmt"

func findDuplicate(nodes []string) []string {
	// States: 0 = unseen, 1 = seen once, 2 = duplicate already reported.
	seen := make(map[string]uint8)
	duplicates := make([]string, 0)
	for _, node := range nodes {
		switch seen[node] {
		case 0:
			seen[node] = 1
		case 1:
			duplicates = append(duplicates, node)
			seen[node] = 2
		}
	}
	return duplicates
}
func main() {
	// Example usage of the findDuplicate function
	nodes := []string{"a", "b", "a", "d", "e", "c", "a", "d", "b", "j"}
	duplicate := findDuplicate(nodes)
	fmt.Println("The duplicate strings are:", duplicate)
}
