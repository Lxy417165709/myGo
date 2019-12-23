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


func getLength(test *ListNode) int{
	length:=0
	for ;test!=nil ;test = test.Next  {
		length++
	}
	return length
}
func rotateRight(head *ListNode, k int) *ListNode {
	length := getLength(head)
	if length==0 {
		return nil
	}else{
		k = k % length
		if k==0 {
			return head
		}
	}


	L := new(ListNode)
	L.Next = head
	a := L.Next
	b := L.Next


	for i:=1;i<=k-1 ;i++  {
		b = b.Next
	}
	for ;b.Next!=nil ;  {
		a = a.Next
		b = b.Next
	}
	b.Next = L.Next

	for ;L.Next.Next!=a ;  {
		L.Next= L.Next.Next
	}
	L.Next.Next = nil
	L.Next = a

	return L.Next

}
func main() {
	l2:=&ListNode{0,nil}
	//l2.Next = &ListNode{1,nil}
	//l2.Next.Next = &ListNode{2,nil}
	x := rotateRight(l2,4)
	//for ;l2!=nil ;l2 = l2.Next  {
	//	fmt.Println(l2.Val)
	//}
	for ;x!=nil ;x = x.Next  {
		fmt.Println(x.Val)
	}
}


/*
	总结
	1. 链表常用双指针法
	2. 注意在有用的代码删除后，记得在适合的位置写上
*/