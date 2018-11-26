package tree

import (
	"fmt"
	"github.com/harrylee2015/harry_tools/tree/queue"
	"github.com/harrylee2015/harry_tools/tree/stack"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//前序遍历
func (root *Node) PreTravesal() {
	if root == nil {
		return
	}
	s := stack.NewStack()
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*Node)
		fmt.Println(cur.Val)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

//中序遍历
func (root *Node) InTravesal() {
	if root == nil {
		return
	}

	s := stack.NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		if s.Empty() {
			break
		}

		cur = s.Pop().(*Node)
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

//后序遍历
func (root *Node) PostTravesal() {
	if root == nil {
		return
	}

	s := stack.NewStack()
	out := stack.NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		out.Push(cur)

		if cur.Left != nil {
			s.Push(cur.Left)
		}

		if cur.Right != nil {
			s.Push(cur.Right)
		}
	}

	for !out.Empty() {
		cur := out.Pop().(*Node)
		fmt.Println(cur.Val)
	}
}

//广度优先遍历
func (root *Node) LevelTravesal() {
	if root == nil {
		return
	}

	linkedList := queue.NewLinkedQueue()
	linkedList.Offer(root)

	for !linkedList.IsEmpty() {
		cur := linkedList.Poll().(*Node)
		fmt.Println(cur.Val)

		if cur.Left != nil {
			linkedList.Offer(cur.Left)
		}

		if cur.Right != nil {
			linkedList.Offer(cur.Right)
		}
	}
}
