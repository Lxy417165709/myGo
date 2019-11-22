package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
     Val int
     Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	L :=&ListNode{-1,head}	//L指向head,做首节点，这是做更方便
	// 初始化尾指针，目标指针
	tail,target:=L,L

	// 尾指针先走n步，题目限制了n<=head节点个数
	for i:=1;i<=n ;i++  {
		tail = tail.Next
	}

	// 尾指针和目标指针一起走，知道尾指针到达链表尾部(此时tail.Next == nil)
	for ;tail.Next!=nil;  {
		target = target.Next
		tail = tail.Next
	}

	// 删除链表倒数第n个链表的操作，要先判断要删除的是否是倒数第一个，因为此时target.Next==nil
	if target.Next!=nil {
		target.Next = target.Next.Next
	}else{
		target.Next = nil
	}
	return L.Next
}

func main() {
	l2:=&ListNode{1,nil}
	l2.Next = &ListNode{2,nil}
	l2.Next.Next = &ListNode{3,nil}
	l2.Next.Next.Next = &ListNode{4,nil}
	l2.Next.Next.Next.Next = &ListNode{5,nil}
	x:=removeNthFromEnd(l2,5)
	for ;x!=nil ;x = x.Next  {
		fmt.Println(x.Val)
	}
}

/*
	总结
	1. 本题我自己采用的是存储思维，遍历一次把链表信息存储起来，再通过下标的方式实现删除
	2. 其实还有另外一种一次遍历方式，等间隔推进法，可以看LeetCode官方,我在上面实现了

*/