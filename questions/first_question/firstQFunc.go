package first_question

import (
	"sort"
	"strings"
)

// SortWordsByA sorts words based on 'a' count and length.
func SortWordsByA(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i != countA_j {
			return countA_i > countA_j
		}

		return len(words[i]) > len(words[j])
	})
	return words
}
