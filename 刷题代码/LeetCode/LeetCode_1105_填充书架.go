package main

import "fmt"


func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
//func minHeightShelves(books [][]int, shelf_width int) int {
//	dp:=make([]int,1)
//	for i:=0;i< len(books);i++  {
//		nowWidth:=0
//		height:=0
//		ans:=1000000000
//		//dp = append(dp, 100000000)
//		for j:=i;j>=0 && nowWidth+books[j][0]<=shelf_width ;j--  {
//			nowWidth = nowWidth+books[j][0]
//			height=max(height,books[j][1])
//			ans =min(ans,dp[j]+height)
//		}
//		dp = append(dp, ans)
//	}
//	fmt.Println(dp)
//	return dp[len(books)]
//
//}

var m map[int]int
func minHeightShelves(books [][]int, shelf_width int) int {
	m=make(map[int]int)
	return solve(len(books)-1,books,shelf_width)
}
func solve(now int,books [][]int, shelf_width int) int{
	if now<0 {
		return 0
	}

	if x,ok :=m[now];ok {
		return x
	}

	ans:=10000000
	height:=0
	width:=0
	for i:=now;i>=0;i--  {
		width +=books[i][0]
		height=max(height,books[i][1])
		if width<=shelf_width {
			ans = min(ans,solve(i-1,books,shelf_width) + height)
		}
	}
	m[now]=ans
	return m[now]
}
func main() {
	fmt.Println(minHeightShelves([][]int{
		//{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2},
	},10))
}
/*
	总结
	1. 这题目不是很懂 ...看了官方解答后自己照着想法写了，虽然中间还是有些不太顺利
	2. dp[i]的含义就是放上这本书的最小层，这个状态可以由它前面的状态得出，dp[i] = min(dp[i],dp[j]+ [j+1,i]的最大高度) 0.0
	3. 官方有用dfs写的 可以学习思路 0.0
	4. 我自己用记忆化搜索做出来了 0.0 可以研究下
	5. 记忆化搜索的记忆注册可以是实体索引 0.0
*/