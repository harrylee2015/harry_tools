package queue

import (
	"container/list"
	"fmt"
	"sync"
)
// 面向接口编程
type Queue interface {
	Offer(e interface{})
	Poll() interface{}
	Clear() bool
	Size() int
	IsEmpty() bool
}

type SliceQueue struct {
	elements []interface{}
	sync.Mutex
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{}
}

func (queue *SliceQueue) Offer(e interface{}) {
	queue.Lock()
	defer queue.Unlock()
	queue.elements = append(queue.elements, e)
}

func (queue *SliceQueue) Poll() interface{} {
	queue.Lock()
	defer queue.Unlock()
	if queue.IsEmpty() {
		fmt.Println("Poll error : queue is Empty")
		return nil
	}
	firstElement := queue.elements[0]
	queue.elements = queue.elements[1:]
	return firstElement
}

func (queue *SliceQueue) Size() int {
	return len(queue.elements)
}

func (queue *SliceQueue) IsEmpty() bool {
	return len(queue.elements) == 0
}

func (queue *SliceQueue) Clear() bool {
	queue.Lock()
	defer queue.Unlock()
	if queue.IsEmpty() {
		fmt.Println("queue is Empty!")
		return false
	}
	for i := 0; i < queue.Size(); i++ {
		queue.elements[i] = nil
	}
	queue.elements = nil
	return true
}

type LinkedQueue struct {
	list *list.List
	sync.Mutex
}

func NewLinkedQueue() *LinkedQueue {
	list := list.New()
	return &LinkedQueue{list: list}
}

func (queue *LinkedQueue) Offer(e interface{}) {
	queue.Lock()
	defer queue.Unlock()
	queue.list.PushBack(e)
}

func (queue *LinkedQueue) Poll() interface{} {
	queue.Lock()
	defer queue.Unlock()
	if queue.IsEmpty() {
		fmt.Println("Poll error : queue is Empty")
		return nil
	}
	if e := queue.list.Front(); e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

func (queue *LinkedQueue) Size() int {
	return queue.list.Len()
}

func (queue *LinkedQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *LinkedQueue) Clear() bool {
	queue.Lock()
	defer queue.Unlock()
	if queue.IsEmpty() {
		fmt.Println("queue is Empty!")
		return false
	}
	for e := queue.list.Front(); e!=nil;e = queue.list.Front() {
		queue.list.Remove(e)
	}
	return true
}
