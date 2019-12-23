package main

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	return solve(root,L,R)
}

func solve(root *TreeNode,L,R int) *TreeNode{
	if root == nil{
		return nil
	}
	// root.Val>R 说明根，右子树都>R
	if root.Val>R{
		return solve(root.Left,L,R)
	}
	// root.Val<L 说明根，左子树都<L
	if root.Val<L{
		return solve(root.Right,L,R)
	}
	// 说明 L <= root.Val <= R 这样左右子树都要递归操作
	root.Left = solve(root.Left,L,R)
	root.Right = solve(root.Right,L,R)
	return root
}

/*
	总结
	1. 刚开始的时候不知道怎么做，然后想到了一个方法，就是把符合条件的加入两个数组，一个前序遍历(当然后序遍历也可以)，
		一个中序遍历，之后再重新建树。但我觉得这个简单的题目不是这样做的，而且我没用到二叉搜索树的条件。
	2. 看了官方题解后，按着它写了一遍代码后自己理解，然后看上面的代码吧
*/
