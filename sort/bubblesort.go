package sort

func BubbleSort(arr []int) {
	n := len(arr)
	for out := n - 1; out >= 1; out-- {
		for in := 0; in < out; in++ {
			if arr[in] > arr[in+1] {
				swap(arr, in, in+1)
			}
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
