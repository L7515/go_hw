package main

import (
	"regexp"
	"sort"
)

type WordCount struct {
	Word  string
	Count int
}

type WordCounts []WordCount

func (wc WordCounts) Len() int { return len(wc) }

func (wc WordCounts) Swap(i, j int) { wc[i], wc[j] = wc[j], wc[i] }

func (wc WordCounts) Less(i, j int) bool {
	if wc[i].Count == wc[j].Count {
		return wc[i].Word < wc[j].Word
	}
	return wc[i].Count > wc[j].Count
}

func structSort(wordToCountMap map[string]int) *WordCounts {
	wordCounts := make(WordCounts, 0, len(wordToCountMap))
	for k, v := range wordToCountMap {
		wordCounts = append(wordCounts, WordCount{k, v})
	}

	sort.Sort(wordCounts)

	return &wordCounts
}

func prepareWords(str string) *[]string {
	regExpoSplitter := regexp.MustCompile(`[^\s,;\n\t]+`)

	words := regExpoSplitter.FindAllString(str, -1)

	return &words
}

func Top10(str string) []string {
	var top10 []string

	if str == "" {
		return top10
	}

	words := prepareWords(str)
	wordsToCountMap := make(map[string]int)

	for _, val := range *words {
		wordsToCountMap[val]++
	}

	sortedWords := *structSort(wordsToCountMap)

	for i := 0; i < 10; i++ {
		top10 = append(top10, sortedWords[i].Word)
	}

	return top10
}
