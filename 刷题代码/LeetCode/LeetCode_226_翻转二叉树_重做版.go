package main


func invertTree(root *TreeNode) *TreeNode {
	solve(root)
	return root
}

func solve(root *TreeNode){
	if root == nil{
		return
	}

	root.Left,root.Right = root.Right,root.Left
	solve(root.Left)
	solve(root.Right)
	return
}

// 层序遍历版
func invertTree(root *TreeNode) *TreeNode {
	// 注意!
	if root == nil{
		return nil
	}
	queue := []*TreeNode{}
	queue = append(queue,root)
	for len(queue)!=0{
		top:=queue[0]
		top.Left,top.Right = top.Right,top.Left
		queue = queue[1:]
		if top.Left!=nil{
			queue =append(queue,top.Left)
		}
		if top.Right!=nil{
			queue =append(queue,top.Right)
		}
	}
	return root
}




/*
	总结
	1. 翻转二叉树思路: 如果树根为nil,则不翻转，否则交换左右子树，并对左右子树进行翻转
	2. 层序遍历版的思路是: 每次把根节点入队，翻转后，加入左右子树节点(前提是非空)。注意初始时要判断根非空
*/