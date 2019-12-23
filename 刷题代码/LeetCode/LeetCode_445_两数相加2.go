package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Stack struct{
	list []int
}

func (this *Stack)pop()int{
	if len(this.list)==0{
		return -1
	}
	ans:=this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	return ans
}
func (this *Stack)push(a int){
	this.list = append(this.list,a)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	A,B:=&Stack{make([]int,0)},&Stack{make([]int,0)}

	for l1!=nil{
		A.push(l1.Val)
		l1 = l1.Next
	}
	for l2!=nil{
		B.push(l2.Val)
		l2 = l2.Next
	}
	ans:=&ListNode{-1,nil}
	carry := 0
	for true{
		sum:=0
		x,y:=A.pop(),B.pop()

		if x == -1 && y== -1{
			break
		}

		if x!=-1{
			sum += x
		}
		if y!=-1{
			sum += y
		}

		sum+=carry
		ans.Val = sum%10
		carry = sum/10

		tmp:=&ListNode{-1,ans}
		ans = tmp
	}
	if carry!=0{
		ans.Val = carry
		return ans
	}else{
		return ans.Next
	}

}

/*
	总结
	1. 这题我第一次做的时候做复杂了，知道是栈，之后看了官方题解后
	   用它这种方法AC，代码比较优美
*/
