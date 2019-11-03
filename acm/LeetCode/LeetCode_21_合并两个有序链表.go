package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head:=&ListNode{-1,nil}
	l3:= head

	for ; l1!=nil && l2!=nil;{
		if l1.Val<l2.Val {
			l3.Next = &ListNode{l1.Val,nil}
			l3 = l3.Next
			l1 = l1.Next
		}else{
			l3.Next = &ListNode{l2.Val,nil}
			l3 = l3.Next
			l2 = l2.Next
		}
	}

	for ;l1!=nil ;l1 = l1.Next  {
		l3.Next = &ListNode{l1.Val,nil}
		l3 = l3.Next
	}
	for ;l2!=nil ;l2 = l2.Next  {
		l3.Next = &ListNode{l2.Val,nil}
		l3 = l3.Next
	}
	return (*head).Next
}

func main() {
	l1:=&ListNode{8,nil}
	l1.Next = &ListNode{8,nil}
	l1.Next.Next = &ListNode{8,nil}

	l2:=&ListNode{1,nil}
	l2.Next = &ListNode{4,nil}
	l2.Next.Next = &ListNode{8888,nil}
	l2.Next.Next.Next = &ListNode{87888,nil}
	l2.Next.Next.Next.Next = &ListNode{8854588,nil}
	x := mergeTwoLists(l1,l2)
	//for ;l2!=nil ;l2 = l2.Next  {
	//	fmt.Println(l2.Val)
	//}
	for ;x!=nil ;x = x.Next  {
		fmt.Println(x.Val)
	}
}

/*
	总结
	1. 好像没什么 0.0..
*/