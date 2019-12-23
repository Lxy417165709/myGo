package 二叉树专题

import "fmt"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val>q.Val {
		p,q = q,p
	}
	return solve(root,p,q)
}

func solve(root,p,q *TreeNode) *TreeNode{
	if root==nil {
		return nil
	}


	if root.Val>=p.Val && root.Val<=q.Val  {
		return root
	}else{
		if root.Val>=q.Val {
			return solve(root.Left,p,q)
		}
		if root.Val<=p.Val {
			return solve(root.Right,p,q)
		}

	}
	return nil
}

func main() {
	//root:=&TreeNode{6,nil,nil}
	//root.Left = &TreeNode{2,nil,nil}
	//root.Left.Left = &TreeNode{0,nil,nil}
	//root.Left.Right = &TreeNode{4,nil,nil}
	//root.Left.Right.Left = &TreeNode{3,nil,nil}
	//root.Left.Right.Right = &TreeNode{5,nil,nil}
	//
	//
	//root.Right = &TreeNode{8,nil,nil}
	//root.Right.Left= &TreeNode{7,nil,nil}
	//root.Right.Right = &TreeNode{9,nil,nil}
	p:=&TreeNode{6,nil,nil}
	q:=&TreeNode{6,nil,nil}
	fmt.Println(lowestCommonAncestor(nil,p,q))
}
/*
	总结
	1. 做该题要理解二叉查找树的特点 --- 左子树的值<=根的值<右子树的值
	2. 这样只要p,q的值在根值两边，那pq公共节点就是根，否则我们根据他们的大小关系选择一边，因为根值和pq的值只有3种关系(注意,我们通过一些处理把p设定为永≤q的值)
	3. 官方也有用迭代完成的,注意迭代不一定都要dfs或bfs 0.0..
*/