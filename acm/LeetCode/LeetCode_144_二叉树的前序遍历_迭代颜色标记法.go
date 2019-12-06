package main

type Node struct{
	tree *TreeNode
	color int
}


func preorderTraversal(root *TreeNode) []int {

	stack:=make([]Node,0)
	stack = append(stack,Node{root,0})  // 0表示没访问
	ans:=make([]int,0)
	for len(stack)!= 0{
		top:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.tree==nil{
			continue
		}
		if top.color == 0{
			// 遍历顺序 top -> top.Left -> top.Right
			// 注意此时根的颜色标记一定要为1
			stack=append(stack,Node{top.tree.Right,0})
			stack=append(stack,Node{top.tree.Left,0})
			stack=append(stack,Node{top.tree,1})
		}else{
			ans=append(ans,top.tree.Val)
		}
	}
	return ans
}

/*
	总结
	1. 这是使用颜色标记法实现的的带前序遍历。
*/
