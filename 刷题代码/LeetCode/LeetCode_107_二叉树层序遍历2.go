package main

import "fmt"

func levelOrderBottom(root *TreeNode) [][]int {
	ans:=[][]int{}
	solve(root,&ans,0)
	for i:=0;i<len(ans)/2;i++{
		ans[i],ans[len(ans)-1-i] = ans[len(ans)-1-i],ans[i]
	}
	return ans
}

func solve(root *TreeNode,ans *[][]int,lay int){
	if root ==nil{
		return
	}
	if len(*ans)==lay{
		*ans = append(*ans,[]int{})
	}
	fmt.Println(root.Val)

	(*ans)[lay] = append((*ans)[lay],root.Val)

	solve(root.Left,ans,lay+1)
	solve(root.Right,ans,lay+1)
}


/*
	总结
	1. 由于迭代层序遍历已经比较会了，所以我采用了递归层序遍历
	2. 注意 solve函数中的ans是 *[][]int，如果是[][]int,那么在append后，ans内部所指向的地址就不是levelOrderBottom
		中ans所指向的地址了。
*/
