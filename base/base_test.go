package base

import (
	"reflect"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestStringType( t *testing.T){
	// str
	str := "hello world!"
	t.Log("type:",reflect.TypeOf(str).Kind())

	p := (*struct {
		ptr uintptr
		len uint8
	})(unsafe.Pointer(&str))
	t.Log(p)
}

func TestSliceType( t *testing.T){

	var slice []int = make([]int,5,10)
	t.Log("type:",reflect.TypeOf(slice).Kind())

	p := (*struct {
		arrayptr uintptr
		len int
		cap int
	})(unsafe.Pointer(&slice))
	t.Log(p)
	p.cap=20
	t.Log(p.cap, cap(slice))
}
type test struct {
	a bool      // 1  byte
	b int8      // 1  byte  填充2个字节
	c int32     // 4
	d int64     // 8 字节，不需要填充
	e string    // 16字节  //整个结构体填充偏移量最大字段或者默认偏移位两者最小值的最小倍数（当前偏移32位,可以整除16，不需要额外填充）
}
func TestBytes(t *testing.T){
	str := "abcd"
    t.Log("rune len:",utf8.RuneCountInString(str))
	t.Log("string size:",unsafe.Sizeof(str),unsafe.Alignof(str))

	t.Log("test size:",unsafe.Sizeof(test{}))

}

