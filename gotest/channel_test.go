package gotest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
  使用for range 优雅地读取channel中得数据
*/
func TestChannel(t *testing.T) {
	channel1 := make(chan int,10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan int, group *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		group.Done()
	}(channel1, &wg)

	for value := range channel1 {
		fmt.Println(value)
	}
	wg.Wait()
	timeout:=time.Duration(-1)
	fmt.Println("timeout:",timeout)
}
