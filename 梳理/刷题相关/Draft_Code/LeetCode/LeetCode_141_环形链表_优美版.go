package main


func hasCycle(head *ListNode) bool {
	slow,fast:=head,head
	for fast!=nil && fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast{
			return true
		}
	}
	return false
}
/*
	总结
	1. 这是在之前快慢指针基础上的优美化版本 0.0
*/
