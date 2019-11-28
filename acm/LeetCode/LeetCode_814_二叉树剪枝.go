package main
func pruneTree(root *TreeNode) *TreeNode {
	flag := solve(root)
	if flag == true{
		return nil
	}
	return root
}

// 返回该树是否不包含1,并会修改该树的结构
func solve(root *TreeNode) bool{
	if root == nil{
		return true
	}
	left:=solve(root.Left)
	right:=solve(root.Right)

	if left==true{
		root.Left = nil
	}
	if right==true{
		root.Right = nil
	}

	return left && right && root.Val == 0
}




/*
	总结
	1. 这题的思路是，我写了solve函数判断一棵树是否不包含1，是则返回true。
		而如果该树为不全1，那么其父节点指向该树的指针将修改为nil,然后递归下去。
	2. 这种写法的话要注意父节点为0时，加入左右子树和父节点都为0，那么应该返回nil，而不是返回父节点

*/

