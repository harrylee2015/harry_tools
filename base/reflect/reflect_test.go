package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type People interface {
	Age() uint8
	Name() string
	Sex() bool
}

type Work interface {
	Work(workName string) string
}
type Employee struct {
	age    uint8  // 1字节
	salary uint64 // 8字节
	name   string //16字节
	sex    bool   // 1字节
	Height uint8
}

func (e *Employee) Age() uint8 {
	return e.age
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) Sex() bool {
	return e.sex
}
func (e *Employee) Work(workName string) string {
	fmt.Println(fmt.Sprintf("%s work is %s", e.name, workName))
	return fmt.Sprintf("%s work is %s", e.name, workName)
}

func TestReflect(t *testing.T) {
	em := &Employee{
		age:    20,
		salary: 0,
		name:   "harrylee",
		sex:    false,
	}
	ty := reflect.TypeOf(*em)
	v := reflect.ValueOf(em)
	t.Log(ty.Kind())
	t.Log(v.Kind())
	// 1.重新转化为原始类型
	//e := v.Interface().(*Employee)
	//t.Log(e)

	//获取结构体字段
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		t.Log(field.Name)
	}

	//修改字段值
	//私有字段只能采用不安全指针编程
	age := (*uint8)(unsafe.Pointer(v.Pointer() + unsafe.Offsetof(em.age)))
	*age = uint8(30)

	//public字段则可以通过反射直接赋值
	if v.Elem().FieldByName("Height").CanSet() {
		v.Elem().FieldByName("Height").Set(reflect.ValueOf(uint8(180)))
	}
	t.Log("em height:", em.Height)
	//3.动态调用结构体方法
	ty = reflect.TypeOf(em) //方法是指针映射
	for i := 0; i < ty.NumMethod(); i++ {
		method := ty.Method(i)
		if md := v.MethodByName(method.Name); md.IsValid() {
			var values []reflect.Value
			if method.Name == "Work" {
				values = md.Call([]reflect.Value{reflect.ValueOf("coding")})
			} else {
				values = md.Call([]reflect.Value{})
			}
			for _, value := range values {
				if va, ok := value.Interface().(string); ok {
					t.Logf("call %s method, return %s", method.Name, va)
					continue
				}
				if va, ok := value.Interface().(bool); ok {
					t.Logf("call %s method, return %v", method.Name, va)
					continue
				}
				if va, ok := value.Interface().(uint8); ok {
					t.Logf("call %s method, return %v", method.Name, va)
				}
			}
		}
	}
}
