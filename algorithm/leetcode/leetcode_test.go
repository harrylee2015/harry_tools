package leetcode

import "testing"

func Test_BinaryGap(t *testing.T) {
	t.Log(BinaryGap(1041))
	t.Log(BinaryGap(32))
	t.Log(BinaryGap(15))
}

func Test_CyclicRotation(t *testing.T) {
	t.Log(CyclicRotation([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
	t.Log(CyclicRotation([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8))
	t.Log(CyclicRotation([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1))
	t.Log(CyclicRotation(nil, 1))
}

func Test_Match(t *testing.T) {
	t.Log(Match([]int{9, 3, 9, 3, 9, 7, 9, 8}))

}

func Test_Counter(t *testing.T) {
	t.Log(Counter(1, []int{1}))

}

func Test_Solution(t *testing.T) {
	// Solution(54321)

	// Solution(1)

	Solution(100010)

	// Solution(100001)
	// Solution(100001)

}
