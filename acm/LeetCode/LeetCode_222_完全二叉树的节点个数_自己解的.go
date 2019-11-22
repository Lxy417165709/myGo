package 二叉树专题


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	ans,_ := solve(root)
	return ans
}
 // 第一个返回值是节点个数，第二个是层数
func solve(root* TreeNode) (int,int){
	if root==nil{
		return 0,0
	}

	numRight,layRight:=solve(root.Right)

	if 1<<uint8(layRight) - 1!=numRight{
		// 非满二叉树
		return numRight + 1<<uint8(layRight),layRight+1
	}else{
		// 满二叉树
		numLeft,layLeft:=solve(root.Left)
		return numLeft+numRight+1,max(layRight,layLeft)+1
	}

}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}


/*
	总结
	1. 根本：判断一棵树是满二叉树
			我这种方法是很笨的，通过全树遍历再和它的层数判断出它是满二叉树

	2. 我又采取了一种方法做，但是没有用到他的思想 0.0 复杂度和第一次做的差不多(可能还更复杂)
	3. 接下来我用官方的思路做   ---  常规解法和击败100%的Java解法
*/