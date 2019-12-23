package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	return solve(root, key)
}

func solve(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 没有优化的
	// root.Left = solve(root.Left,key)
	// root.Right = solve(root.Right,key)

	// 根据性质优化
	if key > root.Val {
		root.Right = solve(root.Right, key)
	} else {
		root.Left = solve(root.Left, key)
	}

	if root.Val == key {
		right := root.Right
		// 右子树为空，直接返回左子树节点
		if right == nil {
			return root.Left
		}
		// 右子树不为空，则找最小节点
		tmp := right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		tmp.Left = root.Left
		return right
	}
	return root
}

/*
	总结
	1. 思路很简单，我是通过后序遍历AC的。
	2. 思路是: 当找到目标节点后，去它的右子树找最小节点(该节点分布在右子树的最左，只要一直left直到left.Next==nil就可以了)
				找到之后，把目标节点左子树拼接到它(上面的最小节点)的左子树中，然后返回目标节点的右子树就可以了。
				注意: 当目标节点右子树为空时，直接返回左子树就可以了。
	3. 我还忘记优化了...我们根据二叉搜索树的性质，很快找到被删除节点。
	4. 这道题找目标节点的左子树的最大节点也是可以的，差不多。
*/
