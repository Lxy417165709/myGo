package main

func detectCycle(head *ListNode) *ListNode {

	slow,fast := head,head
	loopFlag := false
	for fast!=nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast{
			loopFlag = true
			break
		}
	}
	if loopFlag == true{
		first,second := head,slow
		for first!=second{
			first = first.Next
			second = second.Next
		}
		return first

	}else{
		return nil
	}
}

/*
	总结
	1. 思路还是先判断有没有环，有环才能找入口，否则返回nil
*/