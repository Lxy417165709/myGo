package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast,slow:=head,head
	flag:=0
	for ;fast!=nil && slow!=nil ;  {
		slow = slow.Next
		if fast.Next!=nil {
			fast = fast.Next.Next
		}else{
			return nil
		}
		if slow==fast {
			flag=1
			break
		}
	}
	tmp:=head
	// flag==1表示有环
	if flag==1 {
		for ;tmp!=slow ;  {
			tmp = tmp.Next
			slow = slow.Next
		}
	}else{
		tmp = nil
	}

	return tmp
}

func main() {


	fmt.Println(detectCycle(nil))
}
/*
	总结
	1. 该题考查链表是否有环以及环的入口查找
	2. 上面代码可以更简洁
*/