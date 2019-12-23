package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	return solve(root,sum)
}

func solve(root *TreeNode,sum int) [][]int{
	if root==nil{
		return [][]int{}
	}
	if root.Left==nil && root.Right==nil{
		if sum==root.Val{
			return [][]int{{sum}}
		}
		return [][]int{}
	}

	leftSum:=solve(root.Left,sum-root.Val)
	rightSum:=solve(root.Right,sum-root.Val)
	ans:=[][]int{}

	for i:=0;i<len(leftSum);i++{
		tmp:=make([]int,0)
		tmp = append(tmp,root.Val)
		tmp = append(tmp,leftSum[i]...)
		ans = append(ans,tmp)
	}

	for i:=0;i<len(rightSum);i++{
		tmp:=make([]int,0)
		tmp = append(tmp,root.Val)
		tmp = append(tmp,rightSum[i]...)
		ans = append(ans,tmp)
	}
	return ans
}

/*
	总结
	1. 该方法独立性更高，没有使用到全局变量。
*/
