package sync

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type People struct {
	age    uint8  // 1字节
	salary uint64 // 8字节
	name   string //16字节
	sex    bool   // 1字节

}

func TestSyncPool(t *testing.T) {
	t.Log(unsafe.Sizeof(int(0)))
	t.Log(unsafe.Sizeof(People{}))
	pool := &sync.Pool{
		New: func() interface{} {
			return nil
		},
	}
	t.Log(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 3; i++ {
		pool.Put(&People{
			age:    uint8(i),
			salary: 0,
			name:   "",
			sex:    false,
		})
	}

	for i := 0; i < 20; i++ {

		if people, ok := pool.Get().(*People); ok && people != nil {
			go func(p *People,pool *sync.Pool) {
				fmt.Println("age:",p.age)
				pool.Put(p)
			}(people,pool)
			time.Sleep(time.Millisecond*10)
		}
       fmt.Println("num:",i)
	}
	time.Sleep(time.Second)
}
