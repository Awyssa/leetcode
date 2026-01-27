package main

import (
	"sort"
)

func main() {
}

func topKFrequent(nums []int, k int) []int {
	frequentTracker := make(map[int]int)

	for _, num := range nums {
		frequentTracker[num]++
	}

	countedNums := make([]int, 0, len(frequentTracker))

	for key := range frequentTracker {
		countedNums = append(countedNums, key)
	}

	sort.Slice(countedNums, func(i, j int) bool {
		return frequentTracker[countedNums[i]] > frequentTracker[countedNums[j]]
	})

	l := min(k, len(countedNums))

	numsSlice := countedNums[:l]

	return numsSlice
}
