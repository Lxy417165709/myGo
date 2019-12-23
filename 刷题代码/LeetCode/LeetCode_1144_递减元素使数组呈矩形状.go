package main

import "fmt"


func get(nums []int,flag int) int{
	tmp:=make([]int, len(nums))
	copy(tmp,nums)
	ans:=0
	for i:=1;i< len(tmp);i++  {
		if flag==0 {
			if tmp[i-1] >= tmp[i] {
				ans+= tmp[i-1] - tmp[i] + 1
				tmp[i-1] = tmp[i]-1

			}
		}else{
			if tmp[i-1] <= tmp[i] {
				ans+= tmp[i] - tmp[i-1] + 1
				tmp[i] = tmp[i-1]-1
			}
		}
		flag=flag^1
	}
	fmt.Println(tmp)
	return ans
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}

func movesToMakeZigzag(nums []int) int {

	x,y:=get(nums,0),get(nums,1)
	fmt.Println(x,y)
	mn:=min(x,y)


	return mn
}

func main() {
	fmt.Println(movesToMakeZigzag([]int{
		9,
	}))
}
/*
	总结
	1. 我使用的方法是贪心，因为有贪心选择性质 ，比如1 5 6 要变为 ><形的，最好的就是把6变为4，而不要选更小的，因为这样会使代价更大
	2. 具体实现如下， 先判断<><>形的最小代价，然后判断><><形的最小代价，取最小就OK了 0.0.
*/