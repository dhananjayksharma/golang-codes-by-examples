package main

import (
	"fmt"
	"slices"
)

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes) 
	return string(runes)
}

func getAnnagram(list []string) [][]string {
	if len(list) < 2 {
		return nil
	}

	wordList := make(map[string][]string)
	for _, word := range list {
		sortedWord := sortString(word) 
		wordList[sortedWord] = append(wordList[sortedWord], word)
	}

	result := make([][]string, 0, len(wordList))
	for _, words := range wordList {
		if len(words) > 1 {
			result = append(result, words)
		}
	}
	
	return result
}

func main() {
	// Example usage of the anagram function
	list := []string{"cat", "act", "dog", "god", "ate", "tea", "bat", "tab"}
	status:= getAnnagram(list)
	fmt.Printf("Anagram groups found in the list:%q", status)
}
