package main

import (
	"reflect"
	"fmt"
	"strings"
)

func main() {
    fmt.Println(GetNodeName("192.168.0.175:8802"))
    str :="1"
    strs :=strings.SplitN(str,".",2)
    fmt.Println(strs)
	str1 :=strings.SplitAfter(str,".")
	fmt.Println(str1)
   //p := &People{}
   //p.SetChild(p)
   //fmt.Println("reflect value:",p.DriverBase.childValue)
   //fmt.Println("say:",p.Say())
   ////d :=DriverBase{}
   //typ := reflect.TypeOf(p)
   //fmt.Println("reflect type:",typ)
   //fmt.Println("reflect method:",ListMethodByType(typ))
}
func GetNodeName(node string)string{
	return strings.Split(node,":")[0]
}
func ListMethodByType(typ reflect.Type) map[string]reflect.Method {
	methods := make(map[string]reflect.Method)
	fmt.Println("NumMethod:",typ.NumMethod())
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		//mtype := method.Type
		mname := method.Name
		fmt.Println("method Name:",mname)
		// Method must be exported.
		//if method.PkgPath != ""  {
		//	continue
		//}
		methods[mname] = method
	}
	return methods
}
type People struct {
	DriverBase
}
type Driver interface {
	SetName(name string)
	Say()string
}
type DriverBase struct {
	name string
	child  Driver
	childValue reflect.Value
}
func (d *DriverBase) Say()string{
   return "Hello"
}
func (d *DriverBase) SetName(name string) {
	d.name = name
}
func (d *DriverBase) SetChild(e Driver) {
	d.child = e
	d.childValue = reflect.ValueOf(e)
}
func (p *People)Say()string{
	return "people lange"
}