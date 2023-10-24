package slice

import "fmt"

func modifySlice(a []int) {
	a = append(a, 1)
	fmt.Println(a)
}

func modifySlice2(a []int) []int {
	a = append(a, 1)
	fmt.Println(a)
	return a
}
