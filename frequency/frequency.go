package frequency

import (
	"slices"
	"strings"
)

type freqObj struct {
	dict map[string]int
}

func (f freqObj) Complete(prefix string) []string {
	type kv struct {
		Key   string
		Value int
	}

	var matching []kv

	for key, value := range f.dict{
		if strings.HasPrefix(key, prefix) {
			matching = append(matching, kv{key, value})
		}
	}

	slices.SortFunc(matching, func(a, b kv) int {
			return b.Value - a.Value
		})

	results := make([]string, len(matching))
	for i, kv := range matching {
		results[i] = kv.Key
	}

	return results[:10]
}

func New(freqMap map[string]int) (completer freqObj){
	fO := freqObj{
		dict: freqMap,
	}

	return fO
}