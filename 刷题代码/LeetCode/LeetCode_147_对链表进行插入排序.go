package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {

	dummyHead:=&ListNode{-1000000000000,head}
	pre:=dummyHead
	tmp:=head
	for tmp!=nil{

		H:=dummyHead
		nextPoint:=tmp.Next
		flag:=0
		for H.Next!=tmp{
			if tmp.Val>=H.Val && tmp.Val<=H.Next.Val{

				tmp.Next = H.Next
				H.Next = tmp
				pre.Next = nextPoint
				flag=1
				break
			}
			H = H.Next
		}
		if flag==0{
			pre = tmp
		}
		tmp = nextPoint

	}
	return dummyHead.Next
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {

	dummyHead:=&ListNode{-1000000000000,head}

	for head!=nil{
		pre:=dummyHead
		for pre.Next!=nil && pre.Next.Val < head.Val{
			pre = pre.Next
		}
		tmp:=head
		head=head.Next
		tmp.Next = pre.Next
		pre.Next = tmp

	}
	return dummyHead.Next


}
/*
	总结
	1. 这是可以AC的代码，但是不够优美，下面有优美的
*/
