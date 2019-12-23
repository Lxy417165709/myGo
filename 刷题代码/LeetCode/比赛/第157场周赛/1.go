package main

func min(a,b int)int{
	if a>b {
		return b
	}
	return a

}
func minCostToMoveChips(chips []int) int {
	x:=0
	y:=0
	for i:=0;i< len(chips);i++  {
		if chips[i]%2==1 {
			x++
		}else{
			y++
		}
	}
	return min(x,y)
}


func main() {

}
/*
	总结
	1. 这题要思维转换一下，只需要判断数组中返回奇数多还是偶数多，偶数多则返回奇数个数，否则返回偶数个数。
*/