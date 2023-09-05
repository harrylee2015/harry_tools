package leetcode

import "fmt"

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
