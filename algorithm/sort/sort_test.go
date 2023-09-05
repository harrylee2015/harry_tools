package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := QuickSort(arr)

	t.Log(nums)
}

func TestInsertSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := InsertSort(arr)

	t.Log(nums)
}

func TestBubbleSort2(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := BubbleSort2(arr)

	t.Log(nums)
}
func TestSelectionSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := SelectionSort(arr)

	t.Log(nums)
}

func TestShellSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := shellSort(arr)

	t.Log(nums)
}

func TestMergeSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := MergeSort(arr)

	t.Log(nums)
}

func TestHeadSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := heapSort(arr)

	t.Log(nums)
}

func TestCountingSort(t *testing.T) {
	arr := []int{1, 3, 2, 8, 7, 6, 4, 3, 1, 2}

	nums := countingSort(arr)

	t.Log(nums)
}
