package main

import (
	"fmt"
	"sync/atomic"
	//"sync"
	"time"
)
//查找域
var goal int
var count int64

func main() {
	goal = 10000
	c := make(chan int)
    go primetask(c)
	for i := 2; i < goal; i++ {
		c <- i
	}
	time.Sleep(5*time.Second)

}
//类似递归方式，下一个素数是由上一个素数校验后打印出来，依次远远不断，缺点，是每有一个素数，就会有一个 for 循环
//比较占资源
func primetask(c chan int) {
	p := <-c
	if p > goal {
		return
	}
	fmt.Println(p)
	nc := make(chan int)
	fmt.Println("count:",atomic.AddInt64(&count,1))
	go primetask(nc)
	for {
		i := <-c
		if i%p != 0 {
			nc <- i
		}
	}
}
