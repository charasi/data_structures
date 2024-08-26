package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	a := []int{8, -9, 25, 2, 3, 5, 4, 6, 12, 23, 11, -7, -5, -2, -1, 16}
	BubbleSort(a)

}

func TestSelectionSort(t *testing.T) {
	a := []int{8, -9, 25, 2, 3, 5, 4, 6, 12, 23, 11, -7, -5, -2, -1, 16}
	SelectionSort(a)
}

func TestInsertionSort(t *testing.T) {
	a := []int{8, -9, 25, 2, 3, 5, 4, 6, 12, 23, 11, -7, -5, -2, -1, 16}
	InsertionSort(a)
}

func TestMergeSort(t *testing.T) {
	a := []int{8, -9, 25, 2, 3, 5, 4, 6, 12, 23, 11, -7, -5, -2, -1, 16}
	arr := MergeSort(a)
	if len(arr) != len(a) {
		t.Errorf("len(arr)=%d != len(a)=%d", len(a), len(arr))
	}
}
