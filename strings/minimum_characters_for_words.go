package strings

type minimumCharactersForWords struct {
}

func NewMinimumCharactersForWords() minimumCharactersForWords {
	return minimumCharactersForWords{}
}

func (m minimumCharactersForWords) MinimumCharactersForWords(words []string) []string {
	countTotal := make(map[rune]int)
	res := make([]string, 0)

	for _, word := range words {
		count := m.countCharacters(word)
		for char, currentCount := range count {
			countTotal[char] = m.maxVal(countTotal[char], currentCount)
		}
	}

	for char, currentCount := range countTotal {
		res = append(res, m.putChars(char, currentCount)...)
	}

	return res
}

func (minimumCharactersForWords) putChars(char rune, times int) []string {
	res := make([]string, 0)
	for i := 0; i < times; i++ {
		res = append(res, string(char))
	}

	return res
}

func (minimumCharactersForWords) maxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (minimumCharactersForWords) countCharacters(word string) map[rune]int {
	count := make(map[rune]int)

	for _, char := range word {
		count[char]++
	}

	return count
}
