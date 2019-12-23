package main

import "fmt"

func minimumMoves(arr []int) int {
	mp = make(map[[2]int]int)
	return solve(arr,0,len(arr)-1)
}
var mp map[[2]int]int


func solve(arr []int,l int, r int) int{


	if x,ok:=mp[[2]int{l,r}];ok{
		return x
	}
	if l>=r {
		mp[[2]int{l,r}] = 1
		return mp[[2]int{l,r}]
	}

	// 根据题意，由于回文可以删除，于是如果头尾相同，那么我们此时最好就删除了头尾（代价1，但删除了2个元素）
	// 对余下(l+1,r-1)进行遍历就OK了!
	if arr[l]==arr[r]{
		mp[[2]int{l,r}] = solve(arr,l+1,r-1)
		return mp[[2]int{l,r}]
	}	// a Point
	ans:=100000000

	for i:=l;i<=r-1;i++{

		x:=solve(arr,l,i)+solve(arr,i+1,r)
		ans = min(ans,x)
	}
	mp[[2]int{l,r}] = ans
	return mp[[2]int{l,r}]
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

func main() {
	fmt.Println(len([]int{
		1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,1,4,7,8,8,8,8,4,4,4,8,7,12,1,1,3,4,1,5,
	}))
	// 如果把 a Point处删除了，那么答案一定是这个切片的长度（相当于每次只能删除数组中的一个元素）
	// 如果不删，我们就可以处理回文，就可以节省次数（相当于每次可以删除数组头尾两个元素，这两个元素是相同的）
}

/*
	总结
	1. 比赛的时候不知道怎么做，结束后看了官方题解后AC了
	2. AC了后还是觉得有点想不通，于是就在思考了下，想明白了一点，思考过程在上面
*/
