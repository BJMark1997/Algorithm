package main

import (
	"math"
	"strings"
)

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	minDict := len(wordsDict)
	for i := 0; i < len(wordsDict); i++ {
		if strings.Compare(wordsDict[i], word1) == 0 {
			for j := 0; j < len(wordsDict); j++ {
				if strings.Compare(wordsDict[j], word2) == 0 {
					minDict = int(math.Min(float64(minDict), math.Abs(float64(i)-float64(j))))
				}
			}
		}
	}
	return minDict
}

func shortestDistance1(words []string, word1 string, word2 string) int {
	index1, index2, result := -1, -1, len(words)
	for k, v := range words {
		if v == word1 {
			index1 = k
		} else if v == word2 {
			index2 = k
		}
		if index1 != -1 && index2 != -1 {
			result = int(math.Min(float64(result), math.Abs(float64(index1)-float64(index2))))
		}
	}
	return result
}
