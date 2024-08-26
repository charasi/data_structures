package sort

func InsertionSort(arr []int) {
	n := len(arr)
	for out := 1; out < n; out++ {
		in := out
		for in > 0 {
			if arr[in] < arr[in-1] {
				swap(arr, in, in-1)
			}
			in--
		}
	}
}
