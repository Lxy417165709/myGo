package 二叉树专题

import "fmt"

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


var m map[*TreeNode] int
func levelOrder(root *TreeNode) [][]int {
	m := make(map[*TreeNode] int)
	slice:= make([][]int,0)
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	m[root] = 0
	for ; len(queue)!=0;  {
		top:=queue[0]
		if top!=nil {
			// 长度不够，就在加~
			if len(slice)==m[top] {
				slice = append(slice, []int{})
			}
			slice[m[top]] = append(slice[m[top]], top.Val)

			m[top.Left], m[top.Right] = m[top] + 1,m[top] + 1
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
		}
		queue = queue[1:]
	}

	return slice
}


func main() {
	root :=&TreeNode{3,nil,nil}
	//root.Left = &TreeNode{9,nil,nil}
	//root.Left.Left = &TreeNode{10,nil,nil}
	//root.Left.Left.Left = &TreeNode{11,nil,nil}
	//root.Left.Left.Left.Left = &TreeNode{12,nil,nil}
	//root.Left.Left.Left.Left.Left = &TreeNode{15,nil,nil}
	//root.Left.Left.Left.Left.Right = &TreeNode{13,nil,nil}
	//
	//root.Right = &TreeNode{20,nil,nil}
	//root.Right.Left = &TreeNode{15,nil,nil}
	//root.Right.Right = &TreeNode{7,nil,nil}
	fmt.Println(levelOrder(root))
}

/*
	总结
	1. 该题我利用了哈希判断节点是第几层，通过BFS进行遍历加入切片
	2. 官方BFS判断同层时先获取队列长度了，然后确定该队列属于那一层，把队列中该长度的元素再加入切片
	3. 官方DFS把层数作为参数传入~答案的二维数组作为全局变量
*/