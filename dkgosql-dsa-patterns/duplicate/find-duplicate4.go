package main

import "fmt"

func findDuplicate(nodes []string) map[string]int {
	return findDuplicateInWindow(nodes, len(nodes))
}

// findDuplicateInWindow returns each duplicate's maximum count in any window
// containing at most windowSize consecutive elements.
func findDuplicateInWindow(nodes []string, windowSize int) map[string]int {
	seen := make(map[string]int)
	duplicates := make(map[string]int)
	if windowSize <= 1 {
		return duplicates
	}

	// i is the left boundary; j is the right boundary (inclusive).
	for i, j := 0, 0; j < len(nodes); j++ {
		seen[nodes[j]]++

		// Shrink the window from the left when it exceeds windowSize.
		for j-i+1 > windowSize {
			seen[nodes[i]]--
			if seen[nodes[i]] == 0 {
				delete(seen, nodes[i])
			}
			i++
		}

		if count := seen[nodes[j]]; count > 1 && count > duplicates[nodes[j]] {
			duplicates[nodes[j]] = count
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
