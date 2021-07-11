package main

func partitionDisjoint(A []int) int {

	if len(A)==0{
		return 0
	}


	leftMax:=make([]int,len(A))
	rightMin:=make([]int,len(A))
	leftMax[0] = A[0]
	rightMin[len(A)-1] = A[len(A)-1]
	for i:=1;i<len(A);i++{
		leftMax[i] = max(leftMax[i-1],A[i])
	}
	for i:=len(A)-2;i>=0;i--{
		rightMin[i] = min(rightMin[i+1],A[i])
	}

	border:=0
	// 由于题目说一定存在答案
	// 所以 border+1 不会越界，否则就是非法答案了
	// 当然，border<len(A)-1 这样写也可以
	for border<len(A) && leftMax[border]>rightMin[border+1]{
		border++
	}

	return border+1
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}


/*
	总结
	1. 这题思路: 记录[0..i]最大值 以及 [j...n-1]最小值
		然后通过比较就可以得出答案了
*/
