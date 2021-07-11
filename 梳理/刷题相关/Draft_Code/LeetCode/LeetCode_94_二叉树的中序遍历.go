package main

import "fmt"

type Node struct {
	node *TreeNode
	color int
}
func inorderTraversal(root *TreeNode) []int {
	stack:=make([]Node,0)
	stack = append(stack, Node{root,0})
	ans:=make([]int,0)
	for len(stack)!=0{
		top:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node==nil{
			continue
		}
		if top.color==0{
			// 这样表示后序遍历
           	// stack = append(stack, Node{top.node,1})
			// stack = append(stack, Node{top.node.Right,0})
			// stack = append(stack, Node{top.node.Left,0})

			// 这样表示中序遍历
			stack = append(stack, Node{top.node.Right,0})
			stack = append(stack, Node{top.node,1})
			stack = append(stack, Node{top.node.Left,0})

			// 这样表示先序遍历
			// stack = append(stack, Node{top.node.Right,0})
			// stack = append(stack, Node{top.node.Left,0})
			// stack = append(stack, Node{top.node,1})
		}else{
			ans = append(ans, top.node.Val)
		}
	}
	return ans
}
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
	3. 标记是1表示已经加入过队列了，所以可以输出，而加入0的节点表示现在加入，但是还不能输出
*/