package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{-1,nil}
	tmp:=dummyHead
	for l1!=nil && l2!=nil{
		if l1.Val<=l2.Val{
			tmp.Next = l1
			l1 = l1.Next
		}else{
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1!=nil{
		tmp.Next = l1
	}else{
		tmp.Next = l2
	}
	return dummyHead.Next
}

/*
	总结
	1. 这个是挺优雅的迭代版本了，时间复杂度O(n+m),空间复杂度O(1) (因为只有几个指针变量，而链表节点的空间是本来就存在的)
*/