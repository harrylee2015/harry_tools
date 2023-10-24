package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// 二进制间隔
func BinaryGap(n int) int {
	s := fmt.Sprintf("%b", n)
	fmt.Println(s)
	distance, last := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			if last > distance {
				distance = last
			}
			last = 0
			continue
		}
		last++
	}
	return distance
}

// 循环队列
func CyclicRotation(A []int, K int) []int {
	if A == nil {
		return nil
	}
	len := len(A)
	remainder := K % len
	if remainder == 0 {
		return A
	}
	//以空间换时间
	tmp := make([]int, remainder)
	for i := 0; i < remainder; i++ {
		tmp[remainder-1-i] = A[len-1-i]
	}
	//剩余数组前移
	for i := 0; i < len-remainder; i++ {
		A[len-i-1] = A[len-remainder-1-i]
	}
	for i := 0; i < remainder; i++ {
		A[i] = tmp[i]
	}

	return A
}

// func reverse(B []int) []int {
// 	len := len(B)
// 	for i := 0; i < len/2; i++ {
// 		B[len-1-i], B[i] = B[i], B[len-1-i]

// 	}
// 	return B
// }

// 配对
func Match(A []int) int {
	len := len(A)
	if len%2 != 1 {
		panic("error")
	}
	count := make(map[int]int, len)
	for i := 0; i < len; i++ {
		count[A[i]]++
	}
	for K, v := range count {
		if v%2 == 1 {
			return K
		}
	}
	panic("errror")
}

// times
func Times(X, Y, D int) int {
	abs := int(math.Abs(float64(X) - float64(Y)))
	if abs%D != 0 {
		return abs/D + 1
	}
	return abs / D
}

// find
func Find(A []int) int {
	tmp := make([]int, len(A)+2)
	for i := 0; i < len(A); i++ {
		tmp[A[i]] = 1
	}
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == 0 {
			return i
		}
	}
	return 0
}

// ABS
func abs(A []int) int {
	len := len(A)
	abs := -1
	total := 0
	tmp := make([]int, len)
	for i := 0; i < len; i++ {
		total += A[i]
		tmp[i] = total
	}
	for i := 0; i < len-1; i++ {
		left := tmp[i]
		right := total - left
		if abs == -1 {
			abs = int(math.Abs(float64(left) - float64(right)))
		} else if int(math.Abs(float64(left)-float64(right))) < abs {
			abs = int(math.Abs(float64(left) - float64(right)))
		}
	}
	return abs
}

// 青蛙过河
func skip(X int, A []int) int {
	len := len(A)
	if X > len {
		return -1
	}
	tmp := make([]int, X+1)
	for i := 0; i < len; i++ {
		if A[i] <= X && tmp[A[i]-1] == 0 {
			tmp[A[i]-1] = i + 1
		}
		if i >= X-1 {
			flag := false
			for j := 0; j < X; j++ {
				if tmp[j] == 0 {
					flag = true
					break
				}
			}
			if !flag {
				return i
			}
		}

	}
	return -1

}

func isOrderList(A []int) int {
	len := len(A)
	tmp := make([]int, len)
	for i := 0; i < len; i++ {
		if A[i] > len {
			return 0
		}
		if tmp[A[i]-1] >= 1 {
			return 0
		}
		tmp[A[i]-1]++
	}
	return 1
}

// 计数器
func Counter(N int, A []int) []int {
	counters := make([]int, N, N)
	lastReset := -1
	max := 0
	min := 0

	for i := 0; i < len(A); i++ {
		if A[i] == N+1 {
			lastReset = i
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i] >= 1 && A[i] <= N {
			if counters[A[i]-1] < min {
				counters[A[i]-1] = min + 1
			} else {
				counters[A[i]-1]++
			}
			if counters[A[i]-1] > max {
				max = counters[A[i]-1]
			}
		}
		if A[i] == N+1 {
			min = max
		}
		if i == lastReset {
			for j := 0; j < N; j++ {
				counters[j] = min
			}
		}
	}
	return counters
}

func findMissingPositive(A []int) int {
	B := make([]int, len(A))
	j := 0
	for i := 0; i < len(A); i++ {
		if A[i] > 0 {
			B[j] = A[i]
			j++
		}
	}
	if len(B) == 0 {
		return 1
	}
	for i := 0; i < j; i++ {
		value := int(math.Abs(float64(B[i])))
		if value > 0 && value <= j {
			if B[value-1] > 0 {
				B[value-1] = -1 * B[value-1]
			}
		}
	}
	for i := 0; i < j; i++ {
		if B[i] > 0 {
			return i + 1
		}
	}

	return j + 1

}

// 汽车相遇数量
func Count(A []int) int {
	count := 0
	len := len(A)
	left := 0
	for i := 0; i < len; i++ {
		if A[i] == 0 {
			left++
		}
	}
	right := len - left
	rightpass := 0
	for i := 0; i < len; i++ {
		if A[i] == 0 {
			count = count + (right - rightpass)
			if count > 1000000000 {
				return -1
			}
		} else {
			rightpass++
		}
	}

	return count

}

func divisionCount(A, B, K int) int {
	if A > B {
		panic("error")
	}
	valueA := A / K
	if A%K != 0 {
		valueA++
	}

	valueB := B / K

	if valueB >= valueA {
		return valueB - valueA + 1
	}
	return 0
}

// 求最小影响因子
func caculate(S string, P, Q []int) []int {
	if len(P) != len(Q) {
		panic("error")
	}
	strs := strings.Split(S, "")
	slices := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		if strs[i] == "A" {
			slices[i] = 1
			continue
		}
		if strs[i] == "C" {
			slices[i] = 2
			continue
		}
		if strs[i] == "G" {
			slices[i] = 3
			continue
		}
		if strs[i] == "T" {
			slices[i] = 4
			continue
		}

	}

	results := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		slice := slices[P[i] : Q[i]+1]
		min := 0
		for j := 0; j < len(slice); j++ {
			if slice[j] == 1 {
				min = 1
				break
			}
			if slice[j] == 2 {
				min = 2
				continue
			}
			if slice[j] == 3 {
				if min != 0 && min > 3 {
					min = 3
				}
				continue
			}
			if slice[j] == 4 && min == 0 {
				min = 4
			}

		}
		results[i] = min
	}

	return results

}

// func Solution(N int) {
// 	var enable_print int
// 	enable_print = N % 10
// 	for N > 0 {
// 		enable_print = (N + 1) % 10
// 		if N%10 != 0 {
// 			enable_print = 1
// 		}
// 		if enable_print == 1 {
// 			fmt.Print(N % 10)
// 		}
// 		N = N / 10

// 	}
// }

// // 邻接矩阵
// func Solution(A,B []int,N int )int{
// 	[][]int graph := make([][]int,N)
// 	[]int degree = make([]int,N)

// 	for(int[] road:roads){
// 		graph[road[0]][road[1]]++;
// 		graph[road[1]][road[0]]++;
// 		degree[road[0]]++;
// 		degree[road[1]]++;
// 	}
// 	int res = Integer.MIN_VALUE;
// 	for(int i=0;i<n-1;i++){
// 		for(int j=i+1;j<n;j++){
// 			int temp = degree[i]+degree[j]-graph[i][j];
// 			res = Math.max(res,temp);
// 		}
// 	}
// 	return res;
// }
