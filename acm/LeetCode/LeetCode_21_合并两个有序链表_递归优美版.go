package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return solve(l1,l2)
}


func solve(l1 *ListNode, l2 *ListNode) *ListNode{
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}

	if l1.Val<=l2.Val{
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = mergeTwoLists(l2.Next,l1)
		return l2
	}
}
/*
	总结
	1. 这是递归的 0.0 但是如果链表过长，可能会造成栈溢出
*/