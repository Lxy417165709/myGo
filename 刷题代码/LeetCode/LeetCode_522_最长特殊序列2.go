package main

import (
	"fmt"
	//"sort"
)
//type arr []string
//
//func (a arr)Len() int{
//	return len(a)
//}
//func (a arr)Swap(i,j int){
//	a[i],a[j] = a[j],a[i]
//}
//
//func (a arr)Less(i,j int) bool{
//	return len(a[i])>len(a[j])
//}



func findLUSlength(strs []string) int {
	//sort.Sort(arr(strs))

	//var m = make(map[string]int)


	ans:=-1

	for i:=0;i< len(strs);i++  {
		flag:=0
		for t:=0;t<len(strs) ;t++  {
			if i==t {
				continue
			}

			a,b:=strs[i],strs[t]

			if a==get(a,b) {
				flag=1
				break
			}

		}
		if flag==0 {
			if ans<len(strs[i]) {
				ans = len(strs[i])
			}
		}

	}
	return ans

}

// 这是获取公共子序列 ...
func get(a,b string) string{
	var dp[100][100]string


	a="x" + a
	b="y" + b
	//fmt.Println(a,b)
	for i:=1;i<= len(a)-1;i++  {
		for t:=1;t<= len(b)-1;t++ {
			if a[i]==b[t] {
				dp[i][t] = dp[i-1][t-1] + string(a[i])
			}else{
				if len(dp[i-1][t])>len(dp[i][t-1]) {
					dp[i][t] = dp[i-1][t]
				}else{
					dp[i][t] = dp[i][t-1]
				}

			}
		}
	}
	return dp[len(a)-1][len(b)-1]
}

func main() {

	fmt.Println(findLUSlength([]string{
		"aaa", "aaa", "aa",
	}))
}
/*

	总结
	1. 这题想复杂了，暴力比较就能过了,连排序都不用...... 还有就是我们不用求最长公共子序列，只是判断a是否为b的子序列就OK了(n+m复杂度)
	2. 脑子短路 ...
	3. 没有排序的话就要全部遍历完,n^2复杂度，有排序的话可以遇到第一个不是子序列的就返回 0.0 这可以常数级的省时间
*/