package trigram

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var wordBreaks = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type PositionalTrigram struct {
	Trigram string
	Index int
}

func makeTrigrams(input string) []string {
	if len(input) == 0 {
		return nil
	}

	words := wordBreaks.Split(strings.ToLower(input), -1)

	// numTrigrams := 0
	// for _, word := range words {
	// 	(len(word)+3)/3
	// }

	trigrams := []string{}
	for _, word := range words {
		trigrams = append(trigrams, "  "+word[:1])
		if len(word) == 1 {
			break
		}

		trigrams = append(trigrams, " "+word[:2])
		if len(word) == 2 {
			break
		}

		for i := 0; i < len(word)-2; i++ {
			trigrams = append(trigrams, word[i:i+3])
		}
		trigrams = append(trigrams, word[len(word)-2:]+" ")
	}

	return trigrams
}

func Trigrams(input string) []string {
	trigrams := makeTrigrams(input)

	sort.Strings(trigrams)

	j := 1
	for i := 1; i < len(trigrams); i++ {
		if trigrams[i-1] != trigrams[i] {
			trigrams[j] = trigrams[i]
			j++
		}
	}
	return trigrams[:j]
}

func sim(count, len1, len2 int) float32 {
	return float32(count) / float32(len1+len2-count)
}

func Similarity(word1, word2 string) float32 {
	trigram1 := Trigrams(word1)
	trigram2 := Trigrams(word2)

	if len(trigram1) == 0 || len(trigram2) == 0 {
		return 0.0
	}

	fmt.Printf("%#v (%d)\n", trigram1, len(trigram1))
	fmt.Printf("%#v (%d)\n", trigram2, len(trigram2))

	i, j := 0, 0
	count := 0
	for i < len(trigram1) && j < len(trigram2) {
		switch {
		case trigram1[i] < trigram2[j]:
			i++
		case trigram1[i] > trigram2[j]:
			j++
		default:
			i++
			j++
			count++
		}
	}

	fmt.Printf("count: %d\n", count)

	return sim(count, len(trigram1), len(trigram2))
}

func WordSimilarity(needle, haystack string) float32 {
	trgm1 := makeTrigrams(needle)
	trgm2 := makeTrigrams(haystack)

	pt := make([]PositionalTrigram, len(trgm1) + len(trgm2))

	for	i, trgm := range trgm1 {
		pt[i].Trigram = trgm
		pt[i].Index = -1
	}

	for i, trgm := range trgm2 {
		pt[i + len(trgm1)].Trigram = trgm
		pt[i + len(trgm1)].Index = i
	}

	sort.Slice(pt, func(i, j int) bool {
		return pt[i].Trigram < pt[j].Trigram
	})

	j := 0
	ulen := 0
	found := make([]bool, len(pt))
	trg2index := make([]int, len(trgm2))
	for i := 1; i < len(pt); i ++ {
		if pt[i].Trigram == pt[i-1].Trigram {
			if found[j] {
				ulen++
			}
			j++
		}

		if pt[i].Index == -1 {
			found[j] = true
		} else {
			trg2index[pt[i].Index] = j
		}
	}


}
