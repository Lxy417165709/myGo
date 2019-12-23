package main



func removeElements(head *ListNode, val int) *ListNode {



	for head!=nil {
		if head.Val == val{
			head = head.Next
		}else{
			break
		}
	}
	tmp:=head
	if head==nil{
		return nil
	}
	for head.Next!=nil {
		if head.Next.Val == val{
			head.Next = head.Next.Next
		}else{
			head = head.Next
		}

	}
	return tmp
}




// 哑头，代码量减少了！
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{-1,head}
	cur := dummyHead
	for cur.Next != nil{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

// 递归
func removeElements(head *ListNode, val int) *ListNode {
	return solve(head,val)
}

func solve(head *ListNode,val int) *ListNode{
	if head==nil{
		return head
	}
	if head.Val==val{
		return solve(head.Next,val)
	}
	head.Next = solve(head.Next,val)
	return head
}



/*
	总结
	1. 这是很简单的移除链表元素，要注意元素可能重复
	2. 接下来我使用哑头AC吧!
	3. 哑头是一个很好的解决方法，因为它可以让原头节点的处理和其他的节点一样，这样我们就不用单独
		判断头节点了。 (使用哑头一般是涉及到链表的插入删除，因为它们会使链表结构发生变化)
*/