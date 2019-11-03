package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return solve(head)
}

func solve(head *ListNode) *ListNode{
	if head==nil || head.Next==nil {
		return head
	}else{
		x:=solve(head.Next)
		head.Next.Next = head
		head.Next=nil
		return x
	}
}
func main() {
	l2:=&ListNode{1,nil}
	l2.Next = &ListNode{2,nil}
	l2.Next.Next = &ListNode{5,nil}
	x:=reverseList(l2)
	for ;x!=nil ;x = x.Next  {
		fmt.Println(x.Val)
	}
}
/*
	总结
	1. 该题我使用递归解决，想了好久 0.0..
	2. 看了官方的递归版本，我改进了一些
*/