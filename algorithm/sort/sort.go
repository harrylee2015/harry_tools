package main

func main() {
}
//快速排序
func QuickSort(nums []int,left,right int){
	val := nums[(left+right)/2]
	i,j := left,right
	for nums[j] > val {
		j--
	}
	for nums[i] < val {
		i++
	}
	nums[i],nums[j] = nums[j],nums[i]
	i++
	j--
	if i < right {
		QuickSort(nums,i,right)
	}
	if j > left {
		QuickSort(nums,left,j)
	}
}
//冒泡排序
func BubbleSort(nums []int){
	length := len(nums)
	for i:=1;i<length;i++{
		for j:=0;j<length-1;j++{
			if nums[j]>nums[i] {
				nums[j],nums[i] = nums[i],nums[j]
			}
		}
	}
}

//二分查找
//二分查找的基础是先做排序
func BinarySearch(nums []int,left,right,val int)int{
	k := (left+right)/2
	if nums[k]>val {
		return BinarySearch(nums,left,k,val)
	}else if nums[k] < val {
		return BinarySearch(nums,k,right,val)
	}else{
		return k
	}
}

func FibonacciRecursion(n int)int{
	if n==0 {
		return 0
	}else if n==1 {
		return 1
	}else{
		return FibonacciRecursion(n-1)+FibonacciRecursion(n-2)
	}
}
//迭代法
func FibonacciFind(n int)int{
	x,y,fib := 0,1,0
	for i:=0;i<=n;i++{
		if i==0 {
			fib=0
		}else if i== 1{
			fib = x+y
		}else{
			fib=x+y
			x,y = y,fib
		}
	}
	return fib
}
