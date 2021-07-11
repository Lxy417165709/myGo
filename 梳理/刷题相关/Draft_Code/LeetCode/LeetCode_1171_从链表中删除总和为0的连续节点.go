package main

import (
	"fmt"
)

type ListNode struct {
      Val int
      Next *ListNode
  }

func removeZeroSumSublists(head *ListNode) *ListNode {

	return solve(head)
}
func solve(head *ListNode) *ListNode{
	if head==nil {
		return nil
	}
	num:=0
	tmp:=head
	for ;tmp!=nil ;  {
		num+=tmp.Val
		if num==0 {
			return solve(tmp.Next)
		}
		tmp=tmp.Next
	}
	head.Next = solve(head.Next)
	return head
}
func main() {
	head:=&ListNode{1,nil}
	head.Next=&ListNode{0,nil}
	head.Next.Next=&ListNode{-1,nil}
	head.Next.Next.Next=&ListNode{-5,nil}
	head.Next.Next.Next.Next=&ListNode{5,nil}

	x:=removeZeroSumSublists(head)
	for x!=nil{
		fmt.Println(x.Val)
		x=x.Next
	}
	//fmt.Println(removeZeroSumSublists(head))
}
/*
	总结
	1. 做这题的时候，刚看到就知道可以用数组解，利用前缀和.
	2. 除了数组解，我想到了一个错误的，就是边走边拓充的双指针法，  前指针走，后指针保持，直到遇见第一个小于0的数，然后后指针追，但是这样的策略是错误的。
	3. 想到了O（N^2）的解法，就是每次计算从头开始到某位置的值，如果为0，则return solve(tmp.next),否则表示不以该头为开始，改用头的下一个为开始，代码如上
*/