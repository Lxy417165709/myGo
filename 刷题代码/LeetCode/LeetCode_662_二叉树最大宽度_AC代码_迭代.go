package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var mp map[*TreeNode]int    // 数节点到其序号的映射

func widthOfBinaryTree(root *TreeNode) int {

	if root ==nil{
		return 0
	}

	mp = make(map[*TreeNode]int)
	queue := make([]*TreeNode,0)
	queue =append(queue,root)
	mp[root] = 0
	mxWidth:=0
	for len(queue)!=0{
		size:=len(queue)
		mxWidth = max(mxWidth,mp[queue[size-1]] - mp[queue[0]] + 1)
		for i:=0;i<size;i++{
			top:=queue[0]
			queue = queue[1:]
			// 这里可以保证top!=nil,因为我们入队的条件时节点非空
			if top.Left != nil{
				mp[top.Left] = 2*mp[top]
				queue = append(queue,top.Left)
			}
			if top.Right != nil{
				mp[top.Right] = 2*mp[top]+1
				queue = append(queue,top.Right)
			}
		}
	}
	return  mxWidth
}

func max(slice ...int) int{
	if len(slice)==1{
		return slice[0]
	}

	a:=slice[0]
	b:=max(slice[1:]...)
	if a>b{
		return a
	}
	return b
}

/*
	总结
	1. 看了官方题解后恍然大悟，然后通过迭代，节省空间，之后就AC了 0.0
	2. 之前超出空间，是因为我用布尔数组存放树在该位置的节点是否存在，之后再得出它的宽度，
		而这次我直接使用哈希，记录序号，这样就可以节省很多空间，于是就AC了 0.0
*/
