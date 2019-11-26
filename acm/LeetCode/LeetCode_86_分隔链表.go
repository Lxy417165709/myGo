package main

func partition(head *ListNode, x int) *ListNode {

	dummyHead1:=&ListNode{-1,nil}
	tmp1:=dummyHead1
	dummyHead2:=&ListNode{-1,nil}
	tmp2:=dummyHead2

	for head!=nil{

		node:=head
		head=head.Next	// 其实可以不用node,因为此时head在接下来的操作中不会成环

		if node.Val<x{
			tmp1.Next = node
			tmp1 = tmp1.Next
		}else{
			tmp2.Next = node
			tmp2 = tmp2.Next
		}
	}
	tmp1.Next = dummyHead2.Next
	tmp2.Next = nil
	return dummyHead1.Next
}

/*
	总结
	1. 这题我是使用了2个辅助的链表AC的,思路是遇到<x的就加入第一个链表，否则第二个
		最后再把2个链表合并就可以了
*/
