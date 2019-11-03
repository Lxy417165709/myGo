package main

import "fmt"

//type Node struct {
//	node *TreeNode
//	color int
//}
//func inorderTraversal(root *TreeNode) []int {
//	stack:=make([]Node,0)
//	stack = append(stack, Node{root,0})
//	ans:=make([]int,0)
//	for len(stack)!=0{
//		top:=stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		if top.node==nil{
//			continue
//		}
//		if top.color==0{
//			// 可以调整顺序
//			stack = append(stack, Node{root.Right,0})
//			stack = append(stack, Node{root,1})
//			stack = append(stack, Node{root.Left,0})
//		}else{
//			ans = append(ans, top.node.Val)
//		}
//
//	}
//	return ans
//}
func max(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}else{
		x := max(arr[:len(arr)-1]...)
		y := arr[len(arr)-1]
		if x>y{
			return x
		}
		return y
	}
}

func main() {
	fmt.Println(max(1,4,3,2))
}
/*
	总结
	1. 这是采用颜色标记法做出的中序遍历，通过调整顺序，还可以实现前后遍历喔
	2. 第一次做的时候也想到了官方的那样的，但是因为没转过来，直接把右子树节点入栈了...
       然后一直做不出来，之后就去看官方题解了 0.0..
*/