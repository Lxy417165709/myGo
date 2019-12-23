package main

import "fmt"


//var ans int
var ans int
var i int
func shoppingOffers(price []int, special [][]int, needs []int) int {
	ans = 100000000000
	solve(price,special,needs,0)
	fmt.Println(i)
	return ans
}


func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func solve(price []int, special [][]int, needs []int,money int){
	i++
	if isEmpty(needs) {
		ans = min(ans,money)
		return
	}
	for i:=0;i< len(special);i++  {
		gift:=special[i]
		cnt:=gift[len(gift)-1]
		if judge(gift,needs) && isGood(price,gift) {
			tmp:=make([]int, len(needs))
			copy(tmp,needs)
			tmp = down(gift,tmp)
			solve(price,special,tmp,money+cnt)
		}
	}
	tmp:=make([]int, len(needs))
	solve(price,special,tmp,money+get(price,needs))

}

func get(price []int,needs []int) int{
	money:=0
	for i:=0;i< len(needs);i++  {
		money +=needs[i]*price[i]
	}
	return money
}
func isGood(price []int,gift []int) bool{
	money:=0
	for i:=0;i< len(gift)-1;i++  {
		money +=gift[i]*price[i]
	}
	return money>=gift[len(gift)-1]
}

func down(gift []int,needs []int) []int{
	for i:=0;i< len(needs);i++  {
		needs[i] -=gift[i]
	}
	return needs
}
func judge(gift []int,needs []int )bool{
	for i:=0;i< len(needs);i++  {
		if gift[i]>needs[i] {
			return false
		}
	}
	return true
}
func isEmpty(needs []int) bool{
	for i:=0;i< len(needs);i++  {
		if needs[i]!=0 {
			return false

		}
	}
	return true
}

func main() {
	fmt.Println(shoppingOffers([]int{
		2,3,4,
	},[][]int{
		{1,1,0,4,},{2,2,1,9},
	},[]int{
		6,6,6,
	}))
}
/*
	总结
	1. 第一次想到了回溯，然后代码写得太烂了，但是也有想到有进制来刻画状态
	2. 第二次搞贪心，但是做着做着也发现了一些bug
	3. 看了官方题解后，他说是dp 0.0 等等再做 0.0.
	4. 我AC了，使用了回溯 0.0..但是有人说是状态DP，我不太懂
	5. 代码还可优化
	6. 我使用map来记录中间结果，但是好像优化程度不高 0.0.  3个6从 66次减少到38次 0.0..
	7. 优化的代码在官方那里 0.0 我这个是重新粘贴老版本的，为了测试下优化程度
*/