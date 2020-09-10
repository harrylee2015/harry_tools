package channel

import (
	"github.com/Workiva/go-datastructures/queue"
	"github.com/stretchr/testify/assert"
	"runtime"
	"sync"
	"testing"
)

// 原生channel基准测试
func BenchmarkChannel(b *testing.B) {
	ch := make(chan int)
	b.ResetTimer()
	runtime.GOMAXPROCS(2)
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
	close(ch)
}

// ring利用atomic原子锁类似实现了channel的功能
func BenchmarkRingBuffer(b *testing.B) {
	rb := queue.NewRingBuffer(1)
	b.ResetTimer()
	runtime.GOMAXPROCS(2)
	//counter := uint64(0)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		//defer wg.Done()
		for {
			_, err := rb.Get()
			assert.Nil(b, err)
			//if atomic.AddUint64(&counter, 1) == uint64(b.N) {
			//	return
			//}
		}
	}()
	for i := 0; i < b.N; i++ {
		rb.Put(i)
	}

	//wg.Wait()
}
