package main
func jump(nums []int) int {
	reach:=0
	mxPlace:=0
	steps:=0
	for i:=0;i<len(nums)-1;i++{
		mxPlace = max(mxPlace,i+nums[i])
		if i==reach{
			steps++
			reach = mxPlace
		}
	}
	return steps
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}


/*
	总结
	1. 这是使用迭代AC的一个方法，过程如下
		维持3个变量，steps表示到达目的的步数，mxPlace表示当前可到达的最远距离，reach表示可达距离
		在可达距离内，我们可以拓展最远距离，之后如果我们到达了可达距离的最后一个，我们就可以多跳一步，更新最远距离
		循环往复，直到到达目的点
*/