package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

/*
	总结
	1. 这是看了别人代码后，自己写出来的代码，自己之前写的是求出两个节点的前缀后，取其最长公共前缀
		再到二叉树中找该前缀对应的节点，这样虽然也是可以的，但是有点麻烦了。
	2. 注意最近公共祖先节点为nil，表示两个节点都不在这颗树上
	   如果只有一个节点在该树上的话，返回的是该节点
	   如果两个都在该树上的话，那么返回的就是最近公共祖先
	3. 注意这题目适应范围: p、q 为不同节点且均存在于给定的二叉树中。 (自己测试1个节点在二叉树上也可以输出)



*/
