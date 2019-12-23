package main


var leftArr []int   // 存放每一层的最左节点的标号(number)
var maxWidth int    // 最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	leftArr = make([]int,0)
	maxWidth = 0
	solve(root,0,1)
	return maxWidth
}

func solve(root *TreeNode,level int,number int){
	if root == nil{
		return
	}
	if len(leftArr)==level{
		leftArr = append(leftArr,number)
	}

	maxWidth = max(maxWidth,number-leftArr[level]+1)
	solve(root.Left,level+1,number*2)
	solve(root.Right,level+1,number*2+1)
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
/*
	总结
	1. 这是递归写法，还是使用到了标号思想
	2. 这题思路是，存储每一层的最左节点的标号，之后通过先序遍历就可以得到结果了
	3. 这题目的标号没溢出
*/
