package main

import "strconv"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead:=&ListNode{-1,nil}
	tmp :=dummyHead
	for head!=nil{
		if head.Next==nil{
			tmp.Next = head
			tmp = tmp.Next
			break
		}


		pre,next:=head,head.Next
		head = head.Next.Next
		tmp.Next = next
		tmp = tmp.Next
		tmp.Next = pre
		tmp = tmp.Next


		/*
		这样写会死循环,因为head会成环 0.0
		tmp.Next = head.Next
		tmp = tmp.Next
		tmp.Next = head
		tmp = tmp.Next
		head = head.Next.Next
		*/

	}
	tmp.Next = nil

	return dummyHead.Next
}

/*
	总结
	1. 这题我是借助一个辅助链表AC的
	2. 上面有一个注意点，可以模拟它的过程，就可以知道为什么head会成环了
	3. 之前写的是递归解法..
*/
