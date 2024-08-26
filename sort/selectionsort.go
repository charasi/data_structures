package sort

func SelectionSort(arr []int) {
	length := len(arr)
	for out := 0; out < length-1; out++ {
		min := out
		for in := out + 1; in < length; in++ {
			if arr[in] < arr[min] {
				min = in
			}
		}
		swap(arr, out, min)
	}
}
