package main


func findTilt(root *TreeNode) int {
	if root == nil{
		return 0
	}

	stack := []*TreeNode{}
	list := []*TreeNode{}
	stack = append(stack,root)
	for len(stack)!=0{
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		list = append(list,top)

		if top.Right!=nil{
			stack =append(stack,top.Right)
		}
		if top.Left!=nil{
			stack =append(stack,top.Left)
		}
	}



	ans := 0
	for i:=len(list)-1;i>=0;i--{
		node := list[i]
		leftSum,rightSum:=0,0
		if node.Left!=nil{
			leftSum = node.Left.Val
		}
		if node.Right!=nil{
			rightSum = node.Right.Val
		}
		ans += abs(leftSum-rightSum)
		node.Val = node.Val + leftSum + rightSum

	}
	return ans

}



func abs(a int) int{
	if a>0{
		return a
	}
	return -a
}
/*
	总结
	1. 迭代写法类似于中序遍历。每次把左右子树的总和加到根节点，这样根节点就保留了其子树的总和了。

*/
