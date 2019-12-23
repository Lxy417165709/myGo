package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{-1,nil}
	curr :=head
	p,q,carry := l1,l2,0
	a,b := p.Val,q.Val
	for p!=nil && q!=nil{
		num:= a + b + carry
		//fmt.Println(num%10)
		curr.Next = &ListNode{num%10,nil}
		curr = curr.Next
		carry = num /10
		if p.Next == nil && q.Next == nil{
			break
		}
		if p.Next == nil {
			a = 0
		}else{
			p = p.Next
			a = p.Val
		}

		if q.Next == nil {
			b = 0
		}else{
			q = q.Next
			b = q.Val
		}
	}
	if carry == 1{
		curr.Next = &ListNode{1,nil}
	}

	return head.Next
}
func main() {
	l1 := &ListNode{9,nil}

	l2 := &ListNode{9,nil}
	l2.Next = &ListNode{9,nil}
	l2.Next.Next = &ListNode{9,nil}
	fmt.Println(l1,l2)
	fmt.Println(addTwoNumbers(l1,l2))
	fmt.Println(l1,l2)
	head := addTwoNumbers(l1,l2)
	fmt.Println(l1,l2)
	for ; head!=nil;head = head.Next  {
		fmt.Println(head.Val)
	}

}


/*
	总结
	1. 这题考查的类似 链表相加，要注意进位
	2. 注意结构体传递传递的是指针
	3. 可以对指针取指再赋值给x,那x就是指向指针(假如是a)的指针，则a的信息发生改变,x的值不发生改变
*/