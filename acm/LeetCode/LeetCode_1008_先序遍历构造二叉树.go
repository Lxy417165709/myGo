package main

func bstFromPreorder(preorder []int) *TreeNode {
	return solve(preorder)
}

func solve(preorder []int) *TreeNode{
	if len(preorder)==0{
		return nil
	}
	rootVal:=preorder[0]
	index:=0
	for ;index<len(preorder);index++{
		if preorder[index]>rootVal{
			break
		}
	}

	root:=&TreeNode{
		Val:rootVal,
		Left:solve(preorder[1:index]),
		Right:solve(preorder[index:]),
	}
	return root
}

/*
	总结
	1. 该题我是使用递归解决的
	2. 官方有用迭代，利用栈解决的
	3. 我觉得复杂度是NlogN，而不是N...因为在递归中还进行了一次0-n的扫描
*/
