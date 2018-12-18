package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) SetName(name string) {
	p.Name = name
}

func main() {
	p := &Person{"alisa", 10}
	t := reflect.TypeOf(*p) //must be value
	fmt.Println("type.Name:", t.Name())
	v := reflect.ValueOf(p).Elem() // point
	k := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := k.Field(i)
		val := v.Field(i)
		fmt.Println("key.Name", key.Name, "val.Type", val.Type(), "val.Interface", val.Interface())
	}
	//v2 := reflect.Indirect(reflect.ValueOf(p))
	//k2 := v2.Type()
	for i := 0; i < v.NumMethod(); i++ {
		key := k.Method(i)
		val := v.Method(i)
		fmt.Println(key.Name, val.Type(), val.Interface())
	}

	v.FieldByName("Name").Set(reflect.ValueOf("Name"))
	fmt.Println(p.Name)

	name := v.MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(name)

}
