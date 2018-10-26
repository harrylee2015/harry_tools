package queue

import "fmt"

const Max_Queue_Len = 100

type Element interface {
	//TODO 定义数据类型
}

type Queue interface {
	Offer(e Element) //向队列中添加元素
	Poll() Element   //移除队列中最前面的元素
	Clear() bool     //清空队列
	Size() int       //获取队列的元素个数
	IsEmpty() bool   //判断队列是否是空
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

//向队列中添加元素,默认最大队列长度为100
func (entry *sliceEntry) Offer(e Element) error {
	if entry.Size() == Max_Queue_Len {
		return fmt.Errorf("The queue is full,len:%v,Please resend it later!", Max_Queue_Len)
	}
	entry.element = append(entry.element, e)
	return nil
}

//移除队列中最前面的额元素
func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		return nil
	}
	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}
