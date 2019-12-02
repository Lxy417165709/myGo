package main


func reorderList(head *ListNode)  {
	head = solve(head)
}

func solve(head *ListNode) *ListNode{
	// 注意这个递归终止条件
	if head==nil || head.Next==nil || head.Next.Next==nil{
		return head
	}

	tmp := head
	// 为了找到末尾节点的前一个节点
	for tmp.Next.Next!=nil{
		tmp =tmp.Next
	}
	next := head.Next
	head.Next = tmp.Next
	tmp.Next = nil
	head.Next.Next = solve(next)
	return head

}

/*
	总结
	1. 这是递归写法，思路就是每次都拼接首末节点，之后对余下部分进行递归
	2. 递归还可以优化，就是找尾节点其实可以顺带的在翻转时返回
*/
