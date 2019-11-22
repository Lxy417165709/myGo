package 二叉树专题


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func lcaDeepestLeaves(root *TreeNode) *TreeNode {

	_,ans :=solve(root,0)
	return ans
}

func solve(root *TreeNode,lay int) (int,*TreeNode) {
	if root == nil{
		return lay,nil
	}

	lay1,left := solve(root.Left,lay+1)
	lay2,right := solve(root.Right,lay+1)

	var ans *TreeNode = nil

	if lay1==lay2{
		ans = root
	}else{
		if lay1<lay2{
			ans = right
		}else{
			ans = left
		}
	}
	return max(lay1,lay2),ans

}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

// ----  解法2 -- 运用官方的思想

/*
var depthestLay int
var ans *TreeNode
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	depthestLay = 0
	ans = nil
	solve(root,0)
	return ans
}

func solve(root *TreeNode,lay int) int {
	if root == nil{
		depthestLay = max(depthestLay,lay)
		return lay
	}

	lay1 := solve(root.Left,lay+1)
	lay2 := solve(root.Right,lay+1)


	if lay1==lay2 && lay1==depthestLay{
		ans = root
	}
	return max(lay1,lay2)

}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

*/




/*
	总结
	1. 这题刚开始就理解错了，然后做了很多错误的操作...WA了3次
	2. 思考了好久之后才知道怎么解决，
		首先得到左右子树的最深叶子节点的最近公共祖先和层数，
		如果左右层数相等，则根节点是最深叶子节点的最近公共祖先，否则取左右子树之中最深的叶子节点的
		公共祖先。  (函数返回2个值 0.0)
	3. 看了官方后又得到了一些点，最深叶子节点的左右子树层数是最深的(即确定的，记为x)，那么此时我们就可以
	   通过判断左右子树的深度，如果左右子树深度都等于x，则更新最近公共祖先，否则不更新。
	   可看解法2

*/
