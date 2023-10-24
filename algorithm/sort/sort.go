package sort

// // 快速排序
// func QuickSort(nums []int, left, right int) {
// 	val := nums[(left+right)/2]
// 	i, j := left, right
// 	for nums[j] > val {
// 		j--
// 	}
// 	for nums[i] < val {
// 		i++
// 	}
// 	nums[i], nums[j] = nums[j], nums[i]
// 	i++
// 	j--
// 	if i < right {
// 		QuickSort(nums, i, right)
// 	}
// 	if j > left {
// 		QuickSort(nums, left, j)
// 	}
// }

// 快速排序
func QuickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 冒泡排序（交换空间）
func BubbleSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return nums
}

// 优化后的冒泡排序
func BubbleSort2(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

	}
	return nums
}

// 选择排序(交换空间)
func SelectionSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

// 插入排序
func InsertSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		//当前处理索引
		index := i
		//倒序插入到合适位置
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[index] {
				nums[index], nums[j] = nums[j], nums[index]
				index = j
			} else {
				break
			}
		}

	}
	return nums
}

// 希尔排序
func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// 堆排序
func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

// func swap(arr []int, i, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// 计数器/桶排序
func countingSort(arr []int) []int {

	//先找出最大值
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
		}
	}

	bucketLen := maxValue + 1
	// 初始为0的数组
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}

// // 桶排序
// func BucketSort(arr []int) []int {
// 	length := len(arr)
// 	bucket := make(map[int]int, length)

// 	for i := 0; i < length; i++ {
// 		bucket[arr[i]]++
// 	}

// }

// 二分查找
// 二分查找的基础是先做排序
func BinarySearch(nums []int, left, right, val int) int {
	k := (left + right) / 2
	if nums[k] > val {
		return BinarySearch(nums, left, k, val)
	} else if nums[k] < val {
		return BinarySearch(nums, k, right, val)
	} else {
		return k
	}
}

// 斐波那契数列算法
func FibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
	}
}

// 迭代法
func FibonacciFind(n int) int {
	x, y, fib := 0, 1, 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			fib = 0
		} else if i == 1 {
			fib = x + y
		} else {
			fib = x + y
			x, y = y, fib
		}
	}
	return fib
}
