package main
func reverseKGroup(head *ListNode, k int) *ListNode {
	return solve(head,k)
}

func solve(head *ListNode, k int) *ListNode{
	if head==nil{
		return nil
	}
	tmp:=head
	num:=1
	for tmp.Next!=nil && num!=k{
		tmp=tmp.Next
		num++
	}
	// 长度为k说明满足翻转条件
	if num==k {
		pre := head
		tmp = head.Next
		for i:=1;i<k;i++{
			next:=tmp.Next
			tmp.Next = pre
			pre = tmp
			tmp = next
		}
		head.Next = solve(tmp,k)
		return pre

	}else{
		return head
	}

}
func main() {
	
}

/*
	总结
	1. 这题目我是手撕的，手撕熟练了不少
	2. 这题我使用递归+迭代逆转AC了
	3. 翻转链表好玩！

*/