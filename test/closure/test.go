package main

import "fmt"

func f(x int) func(int) int {
	g := func(y int) int {
		return x + y
	}
	// 返回闭包
	return g
}

func main() {
	// 将函数的返回结果"闭包"赋值给变量a
	a := f(3)
	// 调用存储在变量中的闭包函数
	res := a(5)
	fmt.Println(res)
	fmt.Println(a(10))
	// 可以直接调用闭包
	// 因为闭包没有赋值给变量，所以它称为匿名闭包
	fmt.Println(f(5)(5))
	abs := func(a, b int) int {
		return a + b
	}
	fmt.Println("rollbak:",rollbackFunc(3,abs))
}
//回调函数
func rollbackFunc(a int, f func(b, c int) int) int {
	return a + f(1, 2)
}
