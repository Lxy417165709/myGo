package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	queue:=make([]*TreeNode,0)

	level:=0
	queue = append(queue,root)
	ans:=make([][]int,0)
	for len(queue)!=0{
		size:=len(queue)
		ans = append(ans,[]int{})
		ans[level] = make([]int,size)
		for i:=0;i<size;i++{

			top:=queue[0]
			queue = queue[1:]
			ans[level][i] = top.Val
			if top.Left!=nil{
				queue =append(queue,top.Left)
			}
			if top.Right!=nil{
				queue =append(queue,top.Right)
			}

		}
		if level&1==1{
			for i:=0;i<size/2;i++{
				ans[level][i],ans[level][size-1-i] = ans[level][size-1-i],ans[level][i]
			}
		}

		level++
	}
	return ans
}
/*
	总结
	1. 刚刚开始想到了这个方法，就是层次遍历后，偶数层翻转就可以了
	2. 但是觉得这样也太水了吧，于是想找另外的遍历方法，然后想到了栈，但是感觉代码很复杂
	   于是又返回这个了 0.0
*/
