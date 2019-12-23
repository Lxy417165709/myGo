package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head,head
	for ; slow != nil && fast != nil; {
		// 慢指针走一步
		slow = slow.Next

		// 快指针走两步
		if fast.Next !=nil {
			fast = fast.Next.Next
		}else{
			return false
		}

		// 判断是否相等
		if slow==fast {
			return true
		}

	}
	return false


	// 写法
	// 1. 初始化快慢指针的值
	// 2. 循环判断快慢指针的值,慢指针每次走一步，快指针每次走两步
	//			2.1 当快慢指针的值相等则表示有环
	//			2.2 当其中一个等于nil则表示无环
}

func main() {


	fmt.Println(hasCycle(nil))
}

/*
	总结
	1. 该题用到了快慢指针法~

*/
