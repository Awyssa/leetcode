package main

import (
	"reflect"
	"sort"
	"testing"
)

func slicesEqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	return reflect.DeepEqual(aCopy, bCopy)
}

func TestTopKFrequent_Example1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}

	// Uncomment when topKFrequent is implemented
	result := topKFrequent(nums, k)
	if !slicesEqualUnordered(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}

	t.Logf("Test Example 1:")
	t.Logf("  Input: nums = %v, k = %d", nums, k)
	t.Logf("  Expected: %v", expected)
}

func TestTopKFrequent_Example2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}

	// Uncomment when topKFrequent is implemented
	result := topKFrequent(nums, k)
	if !slicesEqualUnordered(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}

	t.Logf("Test Example 2:")
	t.Logf("  Input: nums = %v, k = %d", nums, k)
	t.Logf("  Expected: %v", expected)
}

func TestTopKFrequent_SameFrequency(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2

	// Uncomment when topKFrequent is implemented
	result := topKFrequent(nums, k)
	if len(result) != 2 {
		t.Errorf("Test failed: expected length 2, got %d", len(result))
	}

	t.Logf("Test Same Frequency:")
	t.Logf("  Input: nums = %v, k = %d", nums, k)
	t.Logf("  Expected: any 2 elements")
}

func TestTopKFrequent_LargerArray(t *testing.T) {
	nums := []int{4, 1, -1, 2, -1, 2, 3}
	k := 2
	expected := []int{-1, 2}

	// Uncomment when topKFrequent is implemented
	result := topKFrequent(nums, k)
	if !slicesEqualUnordered(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}

	t.Logf("Test Larger Array:")
	t.Logf("  Input: nums = %v, k = %d", nums, k)
	t.Logf("  Expected: %v", expected)
}
