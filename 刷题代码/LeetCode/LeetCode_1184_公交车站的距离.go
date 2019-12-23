package main

import "fmt"

func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum := make([]int,0)
	sum = append(sum, 0)
	for i:=0;i< len(distance);i++  {
		sum = append(sum,distance[i]+sum[i])
	}
	if destination<start {
		start,destination = destination,start
	}
	x:=sum[destination] - sum[start]
	//y:=sum[start] - sum[0] + sum[len(distance)] - sum[destination]	// 其实不用这样算
		// 由于是环形，y = sum[len(distance)] - x就可以了
	y:=sum[len(distance)]-x
	ans:=min(x,y)
	fmt.Println(x,y)
	return ans

}

func main() {
	fmt.Println(distanceBetweenBusStops([]int{
		1,2,3,4,5,
	},2,0))
}
/*
	总结
	1. 这题目是简单题，我使用了前缀和
	2. 有个地方写复杂了 0.0. 看上面
*/