package main

import (
	"fmt"
	"strconv"
)

//var num int
var m map[int]int
var ans string
func getPermutation(n int, k int) string {
	//num=0
	m=make(map[int]int)
	solve("",n,k)
	return ans
}


func fac(a int) int{
	if a==0 {
		return 1
	}
	return a*fac(a-1)
}

func solve(s string,n int,k int) bool{

	// 这里较原来写法修改了一些 0.0.
	if len(s)==n {
		ans = s
		return true
	}

	for i:=1;i<=n ;i++  {
		if m[i]==1 {
			continue
		}
		// 这部分是剪枝..通过剪枝直接0ms
		str:=s+strconv.Itoa(i)
		// 阶乘还能优化，可以把阶乘值先记录 0.0
		if fac(n-len(str))<k {
			k-=fac(n-len(str))
			continue
		}

		m[i]=1
		if solve(str,n,k) {
			return true
		}
		m[i]=0
	}
	return false
}


func main() {
	fmt.Println(getPermutation(3,3))
}
/*
	总结
	1. 该题可以获得全排列，再取出第k个，但是这样会超时
	2. 我们直接遍历[1,n]当达到第k个时返回就可以了，然后我这样ac了，但是时间还是用了很多
	3. 官方有更好的剪枝的方法 0.0.
	4. 康拓展开可以完美的解决这个问题 0.0..
	5. 看了官方题解后，觉得不能偷懒，自己也开始剪枝，然后0ms通过 0.0..
*/