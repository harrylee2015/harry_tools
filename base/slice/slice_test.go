package slice

import "testing"

func Test_ModifySlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	modifySlice(s)
	t.Log(s)

	t.Log(modifySlice2(s))

	t.Log(1 << 3)
	t.Log((8 * 13 / 16) * 2)

}

func Test_Const(t *testing.T) {
	const (
		A = 1 << iota
		B
		C
		D = iota
	)
	t.Log(A)
	t.Log(B)
	t.Log(C)
	t.Log(D)

}
