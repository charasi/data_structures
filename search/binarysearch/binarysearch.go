package binarysearch

func BinarySearch(a []int, value int) int {
	ub := len(a) - 1
	lb := 0
	for {
		if lb > ub {
			return -1
		}
		mid := lb + (ub-lb)/2
		// return index value if value is in array
		if a[mid] == value {
			return mid
		}
		// if a[mid] > value
		// up = mid - 1
		// lb remains unchanged
		if a[mid] > value {
			ub = mid - 1
		} else {
			// if a[mid] < value
			// lb = mid + 1
			// up remains unchanged
			lb = mid + 1
		}
	}

}
