package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA==nil || headB==nil {
//		return nil
//	}
//	// 1.1
//	lA,lB:=1,1
//	tmpA,tmpB :=headA,headB
//	for ;tmpA.Next!=nil ;  {
//		lA++
//		tmpA = tmpA.Next
//	}
//	for ;tmpB.Next!=nil ;  {
//		lB++
//		tmpB = tmpB.Next
//	}
//	if tmpA!=tmpB {
//		return nil
//	}
//	fmt.Println(lA,lB)
//	if lA>lB {
//		lA,lB=lB,lA
//		headB,headA = headA,headB	//注意不是tmp了
//	}
//	// 1.2
//	for i:=0;i<lB-lA ;i++  {
//		headB = headB.Next
//	}
//	for i:=0;i<lA ;i++  {
//		if headA==headB {
//			return headA
//		}
//		headA = headA.Next
//		headB = headB.Next
//	}
//	return nil
//
//
//
//
//
//
//}

// 最优美的解法 (根本在于消除长度差)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
func main() {
	headA := &ListNode{4, nil}
	headA.Next = &ListNode{1, nil}
	headA.Next.Next = &ListNode{8, nil}
	headA.Next.Next.Next = &ListNode{4, nil}
	headA.Next.Next.Next.Next = &ListNode{5, nil}

	headB := &ListNode{1, nil}
	fmt.Println(getIntersectionNode(headB, headA))
}

/*
注意：
    如果两个链表没有交点，返回 null.
    在返回结果后，两个链表仍须保持原有的结构。
    可假定整个链表结构中没有循环。
    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
		来源：力扣（LeetCode）

	总结
	1. 我的思路是：
		1.1 先获取两链表的长度，获取之后得到他们的尾结点，如果相等表名存在相交，否则不相交
		2.1 假设B链表长度不小于A链表，则让B链表跳 len(A)-len(B)步，之后从这个点开始一起跳，如果相等表明就是相交节点
	2. 官方解答有神级简洁方法，根本目的是消除长度差

*/
