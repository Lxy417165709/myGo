package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// 递归写法 ..
func swapPairs(head *ListNode) *ListNode {
	L:=&ListNode{-1,head}
	return solve(L)


}

// 这个递归需要一个空头结点，[head]->a->b  会变为 [head]->b->a
func solve(head *ListNode) *ListNode{
	if head.Next!=nil && head.Next.Next!=nil{
		T1:=head.Next
		T2:=head.Next.Next
		T3:=solve(T2)
		head.Next = T2
		T2.Next = T1
		T1.Next = T3
	}
	return head.Next
}


func main() {
	x:=swapPairs(nil)
	for ;x!=nil ;x = x.Next  {
		fmt.Println(x.Val)
	}
}
