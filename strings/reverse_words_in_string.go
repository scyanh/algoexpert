package strings

import "strings"

type reverseWordsInString struct {
}

func NewReverseWordsInString() reverseWordsInString {
	return reverseWordsInString{}
}

func (r reverseWordsInString) ReverseWordsInString(str string) string {
	words := make([]string, 0)
	starChar := 0

	for i := range str {
		if str[i] == ' ' {
			words = append(words, str[starChar:i])
			words = append(words, " ")
			starChar = i + 1
		}
	}
	words = append(words, str[starChar:])

	r.reverseArray(words)

	res := strings.Join(words, "")

	return res
}

func (reverseWordsInString) reverseArray(words []string) {
	start := 0
	end := len(words) - 1
	for i := range words {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
		if start >= end {
			break
		}
		start++
		end--
	}
}
