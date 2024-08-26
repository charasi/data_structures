package sort

// MergeSort recursive function for merge sort
func MergeSort(arr []int) []int {
	// base case
	if len(arr) <= 1 {
		return arr
	}

	// split array in half
	mid := len(arr) / 2
	// pass left/right side to merge sort
	// keep passing till base case is reached
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// merge split array
	return merge(left, right)
}

// sorts two array
func merge(left []int, right []int) []int {
	l := len(left)
	r := len(right)
	indexL := 0
	indexR := 0
	m := make([]int, l+r)
	indexM := 0
	//
	for indexL < l && indexR < r {
		if left[indexL] < right[indexR] {
			m[indexM] = left[indexL]
			indexL++
		} else {
			m[indexM] = right[indexR]
			indexR++
		}
		indexM++
	}

	if indexL < l {
		for indexL < l {
			m[indexM] = left[indexL]
			indexL++
			indexM++
		}
	} else {
		for indexR < l {
			m[indexM] = right[indexR]
			indexR++
			indexM++
		}
	}

	return m
}
