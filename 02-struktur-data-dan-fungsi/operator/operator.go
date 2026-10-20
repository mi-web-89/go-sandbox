package operator

import (
	"slices"
	"strings"
)

func CheckDuplicates(items []string) []string {
	duplicates := []string{}
	slices.Sort(items)

	for i := 0; i < len(items); i++ {
		if i > 0 && items[i] == items[i-1] {
			duplicates = append(duplicates, items[i])
		}
	}

	return duplicates
}

func RemoveDuplicates(items []string) []string {
	slices.Sort(items)
	result := slices.Compact(items)
	return result
}

func CheckIntersect(area1, area2 []string) []string {
	intersect := []string{}
	area1 = RemoveDuplicates(area1)
	area2 = RemoveDuplicates(area2)

	for _, item := range area1 {
		if slices.Contains(area2, item) {
			intersect = append(intersect, item)
		}
	}

	return intersect
}

func CalculateFreqs(word string) map[string]int {
	freqs := make(map[string]int)

	for _, item := range strings.Fields(word) {
		freqs[item]++
	}

	return freqs
}
