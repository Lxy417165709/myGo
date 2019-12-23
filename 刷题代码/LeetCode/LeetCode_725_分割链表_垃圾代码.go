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

	section := ceil(length,k)
	ans := []*ListNode{}


	for t:=0;t<length%k;t++{
		ans = append(ans,head)
		for i:=0;i<section-1;i++{
			head=head.Next
		}
		next :=head.Next
		head.Next=nil
		head=next
	}
	for len(ans)!=k{
		ans = append(ans,head)
		if head==nil{
			continue
		}
		cnt:=section-1
		if length%k!=0{
			cnt--
		}
		for i:=0;i<cnt;i++{
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

/*
	总结
	1. 这题AC的思路是，先算出需要多少个超出的，之后再根据条件获取未超出的
	2. 这代码很垃圾...太难看了
	3. 我一定要美化它
*/
