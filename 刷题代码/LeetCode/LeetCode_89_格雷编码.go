package main

import "fmt"

func grayCode(n int) []int {
	return solve(n)
}
func rev(a []int) []int{
	b:=make([]int, len(a))
	copy(b,a)
	for i:=0;i< len(a)/2;i++  {
		b[i],b[len(a)-1-i] = b[len(a)-1-i],b[i]
	}
	return b
}
func judge(a,b int) bool{
	x:=0
	for ;a!=0 ||b!=0 ;  {
		if a&1!=b&1 {
			x++
			if x==2 {
				return false
			}
		}
		a>>=1
		b>>=1
	}
	if x==1 {
		return true
	}else{
		return false
	}

}

// 解法1
//func solve(n int) []int{
//	if n==0 {
//		return []int{0}
//	}else{
//
//		x:=solve(n-1)
//		y:=make([]int,len(x))
//		copy(y,x)
//
//		for i:=0;i< len(x);i++  {
//			y = append(y, x[len(x)-1-i] + 1<<uint(n-1))
//		}
//		return y
//	}
//
//}

// 解法2
func solve(n int) []int{
	slice:=make([]int,0)
	slice = append(slice, 0)
	for i:=1;i<1<<uint(n) ;i++  {
		slice = append(slice, i^(i>>1))
	}


	return slice
}

// 错误的
//func solve(l int,r int) []int {
//	var ans []int
//	if l==r {
//		return []int{l}
//	}else{
//		mid := (l+r)/2
//		x:=solve(l,mid)
//		y:=solve(mid+1,r)
//		if len(x)==0{
//			return y
//		}else{
//			if len(y)==0 {
//				return x
//			}
//		}
//		bg1:=x[0]
//		end1:=x[len(x)-1]
//		bg2:=y[0]
//		end2:=y[len(y)-1]
//		if judge(end1,bg2) {
//			ans = append(ans, x...)
//			ans = append(ans, y...)
//		}else{
//			if judge(end2,bg1) {
//				ans = append(ans, y...)
//				ans = append(ans, x...)
//			}else{
//				if judge(bg1,bg2) {
//					ans = append(ans, rev(x)...)
//					ans = append(ans, y...)
//				}else{
//					if judge(end1,end2){
//						ans = append(ans, x...)
//						ans = append(ans, rev(y)...)
//					}
//				}
//			}
//		}
//
//		return ans
//	}
//}
func main() {
	fmt.Println(grayCode(0))
}
/*
	总结
	1. 第一次做的时候用分治法，但是做到一半发现存在逻辑错误 0.0 比如x,y无法合并的情况
	2. 看官方答案的时候，发现可以用位运算做 0.0..思路也知道怎么样了，等等就AC他，先玩一盘 0.0
	3. 刚刚开始做的时候没有思路..我是看官方题解的 0.0.其实应该可以想到的呀 0.0..
	4. 经验：对于数字的话，可以考虑它的*2倍，比如1 2 4 8
*/