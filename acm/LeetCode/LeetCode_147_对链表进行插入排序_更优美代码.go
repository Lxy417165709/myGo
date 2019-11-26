package main
func insertionSortList(head *ListNode) *ListNode {

	dummyHead:=&ListNode{-1000000000000,head}

	for head!=nil && head.Next!=nil{

		// 表示当前有序，则选择下一个排
		if head.Val<=head.Next.Val{
			head = head.Next
			continue
		}

		pre:=dummyHead
		for pre.Next!=nil && pre.Next.Val < head.Next.Val{
			pre = pre.Next
		}
		cur:=head.Next
		head.Next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur

	}
	return dummyHead.Next

}

/*
	总结
	1. 这是一个更优美的代码
	2. 5 6 7 2 x，head.Next指向2时 排序后是 2 5 6 7 x，此时head.Next还是指向2，
		所以上面有
		if head.Val<=head.Next.Val{
			head = head.Next
			continue
		}
		当然，这一步也是为了去除有序链表的排序
*/
