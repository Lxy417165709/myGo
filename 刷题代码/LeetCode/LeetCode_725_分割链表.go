package main
func splitListToParts(root *ListNode, k int) []*ListNode {

	tmp := root
	length:=0
	// 获取长度
	for tmp!=nil{
		length++
		tmp=tmp.Next
	}

	head:=root
	ans := []*ListNode{}

	// 多出一个节点的链表
	for t:=0;t<length%k;t++{
		ans = append(ans,head)
		for i:=0;i<ceil(length,k)-1;i++{
			head=head.Next
		}
		next :=head.Next
		head.Next=nil
		head=next
	}
	// 没有多出节点的链表
	for len(ans)!=k{
		ans = append(ans,head)
		if head==nil{
			continue
		}
		for i:=0;i<floor(length,k)-1;i++{
			head=head.Next
		}
		next :=head.Next
		head.Next=nil
		head=next
	}
	return ans
}

func ceil(a,b int) int{
	if a%b==0{
		return a/b
	}else{
		return a/b+1
	}
}
func floor(a,b int) int{
	return a/b
}
/*
	总结
	1. 这是优美化过的代码..但是看起来还是有些不太好
*/
