package main

func reorderList(head *ListNode)  {
	tmp := head
	length:=0
	for tmp!=nil{
		length++
		tmp = tmp.Next
	}
	if length<=2{
		return
	}
	solve(head,length)

}

func solve(head *ListNode,length int) *ListNode{
	if length==1{
		tail:=head.Next
		head.Next = nil
		return tail
	}
	if length==2{
		tail:=head.Next.Next
		head.Next.Next = nil
		return tail
	}

	tail := solve(head.Next,length-2)
	next := head.Next
	head.Next = tail
	outTail := tail.Next
	tail.Next = next
	return outTail


	// ListNode tail = reorderListHelper(head.next, len - 2);
	// ListNode subHead = head.next;//中间链表的头结点
	// head.next = tail;
	// ListNode outTail = tail.next;  //上一层 head 对应的 tail
	// tail.next = subHead;
	// return outTail;

}

/*
	总结
	1. 这是借鉴了leetcode大佬的思路之后写出的代码，很厉害的一个思路
	2. 思路是 翻转中间链表   1(头)->中间->5(尾)，之后通过一些关系就可以得到尾节点，再返回这个尾节点
		这样我们就不用每次都扫描一次链表来寻找尾节点了
*/

