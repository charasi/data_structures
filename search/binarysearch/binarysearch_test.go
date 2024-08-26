package binarysearch

import "testing"

func TestBinarySearch_SearchFive(t *testing.T) {
	a := []int{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	i := BinarySearch(a, 5)
	if i != 13 {
		t.Error("Expected 13, got ", i)
	}

}
