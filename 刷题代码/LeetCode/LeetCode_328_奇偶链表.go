package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	return solve(head)
}

func solve(head *ListNode) *ListNode{
	if head==nil ||head.Next==nil {
		return head
	}
	A:=head
	B:=head.Next
	C:=B	// B头
	tmp:=B.Next
	for i:=1;tmp!=nil ;i++  {
		if i%2==1 {
			A.Next = tmp
			A = A.Next
		}else{
			B.Next = tmp
			B = B.Next
		}
		tmp = tmp.Next
	}
	B.Next=nil	// 重点，否则会死循环，解决总结的第一条
	A.Next = C

	return head
}
func main() {
	head:=&ListNode{1,nil}
	head.Next=&ListNode{2,nil}
	//head.Next.Next=&ListNode{3,nil}
	//head.Next.Next.Next=&ListNode{4,nil}
	//head.Next.Next.Next.Next=&ListNode{5,nil}
	//tmp:=head
	//for ;tmp!=nil ;  {
	//
	//	fmt.Print(tmp.Val," ")
	//	tmp=tmp.Next
	//}
	//fmt.Println()
	//fmt.Println(solve(nil))
	A:=head
	A.Next=nil
	fmt.Println(A)
	tmp:=head
	for ;tmp!=nil ;  {
		fmt.Print(tmp.Val," ")
		tmp=tmp.Next
	}
	//fmt.Println()


}
/*
	总结
	1. 注意链表的结构，链表并非一个点，而是一个串 0.0..
		1->2->3->4->5	A=head时，B=3->4->5时，A.next=B,得到的是1->3->4->5，而不是1->3
	2. 第一次用递归写，发现一直死循环，之后换迭代，也死循环，之后模拟了一下才发现了问题(就是上面的问题)
	3. A=head,则A.next的修改对head有效,A自身的修改对head无效,可以从内存分析得出
	4. 官方的解答用了类似双指针法，我这个还多出了个tmp指针 0.0..
	5. 解决链表问题可以先在纸上画出结构分析 0.0 注意A=head时，A和head信息相同，地址不同 0.0.
*/