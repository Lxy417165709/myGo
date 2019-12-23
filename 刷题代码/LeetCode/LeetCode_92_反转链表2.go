package main


func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil{
		return nil
	}
	dummyHead:=&ListNode{-1,head}

	first := dummyHead

	for i:=0;i<m-1;i++{
		first = first.Next
	}
	second := first
	for i:=0;i<=n-m;i++{
		second = second.Next
	}
	next := second.Next
	second.Next = nil

	H,T := reverse(first.Next)
	first.Next = H
	T.Next = next

	return dummyHead.Next

}

// 翻转链表，并返回翻转后的头尾节点
func reverse(head *ListNode) (H *ListNode,T *ListNode){
	T = head
	var pre,next *ListNode = nil,nil
	for head!=nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	H = pre
	return H,T
}
/*
	总结
	1. 思路是: 对[m,n]链表执行翻转，之后再进行拼接
	2. 这代码还是有些麻烦了，最好的话是一次扫描，然后在扫描的过程中就可以实现翻转了
	3. dummyHead节点很好用!
*/