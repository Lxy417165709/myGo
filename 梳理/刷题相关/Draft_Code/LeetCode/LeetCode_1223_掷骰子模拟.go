package main
var mod = 1000000007
func dieSimulator(n int, rollMax []int) int {
	mp = make(map[[7]int]int)
	return solve(n,rollMax,rollMax,-1)
}

var mp map[[7]int]int
func solve(n int,base []int,now []int,pre int) int{


	if n== 1{
		return count(now)
	}
	arr:=[7]int{}
	for i:=0;i<6;i++{
		arr[i] = now[i]
	}
	arr[6] = n
	if x,ok:=mp[arr];ok{
		return x
	}


	ans:=0
	tmp:=make([]int,6)
	for i:=0;i<6;i++{
		if now[i]==0{
			continue
		}
		// 随数进行相应变化
		if i!=pre{
			copy(tmp,base)
			tmp[i]--
		}else{
			now[i]--
			copy(tmp,now)
		}
		ans+=solve(n-1,base,tmp,i)%mod
	}
	mp[arr] = ans%mod

	return mp[arr]

}

// 计算切片非0元素个数
func count(arr []int) int{
	ans:=0
	for i:=0;i<len(arr);i++{
		if arr[i]==0{
			continue
		}
		ans++
	}
	return ans
}
func main() {

}

/*
	总结
	1. 这题我是使用递归实现的，思路是每次在数组中取一个数，之后递归剩下的。
	   注意：为了满足要求，数组要随选择的数做出相应变化 0.0
*/

