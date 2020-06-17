package gotest

import (
	"fmt"
	sync2 "sync"
	"testing"
	"time"
)
/**
指针传递
 */
type A struct {
	a int
}
func TestPoint(t *testing.T){
    a :=&A{10}
    var group  sync2.WaitGroup
	group.Add(2)
    go func(P *A,group *sync2.WaitGroup) {
       for i:=0;i<=10;i++  {
		   <-time.After(time.Second)
		   fmt.Println("get",P.a)
	   }
		group.Done()
	}(a,&group)
	go func(P *A,group *sync2.WaitGroup) {
		for i:=0;i<=10;i++ {
			<-time.After(time.Second)
			P.a=i
			fmt.Println("set:",P.a)
		}
		group.Done()
	}(a,&group)
    group.Wait()
}

