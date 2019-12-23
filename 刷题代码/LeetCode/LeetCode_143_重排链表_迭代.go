package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head==nil{
		return
	}
	mid:= getMid(head)
	first,second:=head,reverse(mid.Next)
	mid.Next = nil
	for second!=nil{
		firstNext := first.Next
		secondNext := second.Next
		second.Next = first.Next
		first.Next = second
		first = firstNext
		second = secondNext
	}
}

func getMid(head *ListNode) *ListNode{
	slow,fast:=head,head
	for fast!=nil && fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode{
	var pre,next *ListNode = nil,nil
	for head!=nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre

}
/*
	总结
	1. 迭代思路是: 将链表后半部分进行翻转，这样就有2个链表了，对这两个链表进行合并就可以了
					合并是交替合并， 1->2->3   4->5->6  合并后为 1->4->2->5->3->6
*/