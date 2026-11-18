package main

import (
	"fmt"
	"slices"
)

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	fmt.Println("Sorted runes:", runes)
	return string(runes)
}

func getAnnagram(list []string) bool {
	if len(list) < 2 {
		return false
	}

	wordList := make(map[string]struct{}) 
	for _, word := range list {
		sortedWord := sortString(word)
		fmt.Println("Sorted word:", sortedWord)
		if _, exists := wordList[sortedWord]; exists {
			return true
		}
		wordList[sortedWord] = struct{}{}
	}
	return false
}
func main() {
	// Example usage of the anagram function
	list := []string{"cat", "act"}
	status:= getAnnagram(list)
	if status {
		fmt.Println("Anagrams found in the list.")
	} else {
		fmt.Println("No anagrams found in the list.")
	}
}
