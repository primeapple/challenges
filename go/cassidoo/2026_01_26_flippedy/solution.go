package flippedy

import (
	"cassidoo/utils"
	"slices"
	"strings"
)

func Flippedy(wordsStr string) string {
	words := strings.Split(wordsStr, " ")
	if len(words) <= 1 {
		return wordsStr
	}

	numOfVowels := countVowels(words[0])

	mappedWords := make([]string, len(words))
	mappedWords[0] = words[0]

	for i:= 1; i<len(words); i++ {
		word := words[i]
		if countVowels(word) == numOfVowels {
			reversed := []rune(word)
			slices.Reverse(reversed)
			mappedWords[i] = string(reversed)
		} else  {
			mappedWords[i] = words[i]
		}
	}

	return strings.Join(mappedWords, " ")
}

func countVowels(word string) int {
	return utils.Reduce([]rune(word), 0, func(char rune, acc int) int {
		var newAcc int
		switch char {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			newAcc = acc + 1
		}

		return newAcc
	})
}
