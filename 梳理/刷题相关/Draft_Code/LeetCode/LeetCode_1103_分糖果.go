package main

import (
	"fmt"
	"math"
)

func distributeCandies(candies int, num_people int) []int {

	time:=int(math.Sqrt(float64(2*candies)+0.25)-0.5)
	ans:=make([]int,num_people)
	i:=1
	for ;i<=time ;i++  {
		ans[(i-1)%num_people]+=i
		candies-=i
	}
	ans[(i-1)%num_people]+=candies
	return ans

}
func main() {
	fmt.Println(distributeCandies(1000000000,1))
}
/*
	总结
	1. 这题我是用模拟做出来的0.0，但我肯定这题有公式
*/