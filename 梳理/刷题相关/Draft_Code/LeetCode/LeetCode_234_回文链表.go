package main


func isPalindrome(head *ListNode) bool {
	length:=0
	// 计算链表长度
	tmp:=head
	for tmp!=nil{
		tmp = tmp.Next
		length++
	}

	var pre,next *ListNode
	times:=0
	for times<length/2{
		next=head.Next
		head.Next=pre
		pre = head
		head=next
		times++
	}

	// 长度为奇数时，要后移一位
	if length&1==1{
		head=head.Next
	}
	for true{
		if pre==nil && head == nil{
			return true
		}
		if pre==nil || head == nil{
			return false
		}
		if head.Val != pre.Val{
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}
/*
	总结
	1. 刚刚开始没思路，只想到了空间On的算法
	2. 在舍友提醒下，然后用反转链表，但是第一次是反转整体，但是这样是不可以的，因为会丢失原有的指针
		而如果要新创建一个链表，那么空间还是O(n)
	3. 于是看了下官方题解，发现它是反转一半..然后比较
	4. 我的话是反转前一半，之后比较。官方的大多是反转后一半..
*/