package unsafe

import (
	"testing"
	"unsafe"
)

type People struct {
	age    uint8  // 1字节
	salary uint64 // 8字节
	name   string //16字节
	sex    bool   // 1字节

}

func TestPointer(t *testing.T) {
	people := &People{
		age:    10,
		salary: 1000,
		name:   "XIAOMING",
		sex:    false,
	}
	t.Log(unsafe.Offsetof(people.age),unsafe.Offsetof(people.name))
	t.Log(unsafe.Sizeof(*people),unsafe.Alignof(*people))
	p := unsafe.Pointer(people)
	pointerAge := (*uint8)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(people.age)))
	*pointerAge = uint8(30)
	pointerName := (*string)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(people.name)))
	*pointerName = "harrylee"
	t.Log(people)

}
