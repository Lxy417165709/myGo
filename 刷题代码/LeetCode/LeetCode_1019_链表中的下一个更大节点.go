package main

func nextLargerNodes(head *ListNode) []int {
	var pre,next *ListNode = nil,nil
	for head != nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	ans:=[]int{}
	stack:=[]int{}

	for pre!=nil{

		// 严格递减栈 (不存在相等的 即 7 3 3 2 这种不允许)
		// ---
		for len(stack)!=0 && pre.Val>=stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
		}
		if len(stack)==0{
			ans = append(ans,0)
		}else{
			ans = append(ans,stack[len(stack)-1])
		}
		stack = append(stack,pre.Val)
		// ---
		pre = pre.Next
	}

	for i:=0;i<len(ans)/2;i++{
		ans[i],ans[len(ans)-1-i] = ans[len(ans)-1-i],ans[i]
	}
	return ans

}
/*
	总结
	1. 这题我是使用反转链表后，再用递减栈求出ans切片，最后对ans执行反转就可以了。
*/
