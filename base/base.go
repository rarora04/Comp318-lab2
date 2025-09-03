package base

import (
	"maps"
	"slices"
	"strings"
)

type baseObj struct {
	words []string
}

func (f baseObj) Complete(prefix string) []string {
	tenFirst := make([]string, 0, 10)

	for i := 0; i < len(f.words); i++ {
		word := f.words[i]
		if strings.HasPrefix(word, prefix) {
			tenFirst = append(tenFirst, word)
		}
	}

	return tenFirst[:10]
}

func New(freqMap map[string]int) (completer baseObj) {
	sortedWords := slices.Sorted(maps.Keys(freqMap))

	bO := baseObj{
		words: sortedWords,
	}

	return bO
}
