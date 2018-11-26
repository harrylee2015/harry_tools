package queue
import (
	"fmt"
	"testing"
)
func init(){

}
func TestLinkedQueue_Clear(t *testing.T) {
	q:=NewLinkedQueue()
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Offer(4)
	q.Offer(5)
	fmt.Println("len:",q.Size())
	q.Clear()
	if !q.IsEmpty(){
		t.Errorf("Expected the size of queue to be %d,but instead got %d",0,q.Size())
	}
}

func TestLinkedQueue_Poll(t *testing.T) {
	q:=NewLinkedQueue()
	q.Offer(1)
	fmt.Println("len:",q.Size())
	fmt.Println(q.Poll())
	q.Clear()
	if !q.IsEmpty(){
		t.Errorf("Expected the size of queue to be %d,but instead got %d",0,q.Size())
	}
}