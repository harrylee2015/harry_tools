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
/*
  struct string{
    uint8 *str;
    int len;
  }
  struct []uint8{
    uint8 *array;
    int len;
    int cap;
  }
  uintptr是golang的内置类型，是能存储指针的整型，uintptr的底层类型是int，它和unsafe.Pointer可相互转换。
  但是转换后的string与[]byte共享底层空间，如果修改了[]byte那么string的值也会改变，就违背了string应该是只读的规范了，可能会造成难以预期的影响。
*/
func str2byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0],x[1],x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}
func byte2str(b []byte) string{
	return *(*string)(unsafe.Pointer(&b))
}
func byteStr(b []byte)string{
	return string(b)
}
func strbyte(s string)[]byte{
	return []byte(s)
}
func BenchmarkStr2byte(b *testing.B){
	str :="harrylee 2015  boy";
	b.ResetTimer()
	for i:=0;i<b.N;i++{
        str2byte(str);
	}
}
func BenchmarkStrbyte(b *testing.B){
	str :="harrylee 2015  boy";
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		strbyte(str)
	}
}

func BenchmarkByteStr(b *testing.B){
	bytes :=[]byte("harrylee 2015  boy")
	b.StartTimer();
	for i:=0;i<b.N;i++{
		byteStr(bytes)
	}
	b.StopTimer()
}
func BenchmarkByteStr2(b *testing.B){
	bytes :=[]byte("harrylee 2015  boy")
	b.StartTimer();
	for i:=0;i<b.N;i++{
		byte2str(bytes)
	}
	b.StopTimer()
}