package flippedy

import "strings"

func FlippedyAI(s string) string {
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return ""
	}
	first := words[0]
	vowelCount := countVowelsAI(first)
	result := []string{first}
	for _, word := range words[1:] {
		if countVowelsAI(word) == vowelCount {
			result = append(result, reverseAI(word))
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func countVowelsAI(word string) int {
	count := 0
	for _, r := range word {
		if strings.ContainsRune("aeiou", r) {
			count++
		}
	}
	return count
}

func reverseAI(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
