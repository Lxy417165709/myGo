package main

import "fmt"

//func longestWPI(hours []int) int {
//	ans=0
//
//	solve(hours,0, len(hours)-1)
//
//	return ans
//
//}
//var ans int
//func max(a,b int)int{
//	if a>b {
//		return a
//	}
//	return b
//}
//// 错误分治， 比如，8,10,6,16,5,
//func solve(hours []int,l int,r int) int{
//	if l==r {
//		if hours[l]<=8 {
//			return -1
//		}
//		ans = max(ans,1)
//		return 1
//	}
//	mid:=(l+r)/2
//	left:=solve(hours,l,mid)
//	right:=solve(hours,mid+1,r)
//	//fmt.Println(l,mid,left)
//	//fmt.Println(mid+1,r,right,"----")
//
//	ta,tb:=left,right
//	for i:=mid+1;i< len(hours);i++  {
//		if hours[i]>8 {
//			ta++
//		}else{
//			ta--
//		}
//		if ta>0 {
//			ans = max(ans,i-l+1)
//		}
//	}
//	for i:=mid;i>= l;i-- {
//		if hours[i] > 8 {
//			tb++
//		} else {
//			tb--
//		}
//		if tb > 0 {
//			ans = max(ans, r-i+1)
//		}
//	}
//
//	return left+right
//
//}

// 自己写的ac
//func longestWPI(hours []int) int {
//	m:=make(map[int]int)
//	m[0] = 0
//	tmp:=0
//	ans:=0
//	for i:=1;i<= len(hours);i++ {
//		if hours[i-1] > 8 {
//			tmp = tmp + 1
//		} else {
//			tmp = tmp - 1
//
//		}
//		if tmp>0 {
//			ans = max(ans, i)
//		}else{
//			if x,ok:=m[tmp-1];ok {
//				ans = max(ans,i-x)
//			}else {
//				// 这里很重要...  标注1
//				if _,ok:=m[tmp];!ok{
//					m[tmp] = i
//				}
//			}
//		}
//
//	}
//	return ans
//}

func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}
// 单调栈
func longestWPI(hours []int) int {
	arr:=make([]int,1)
	for i:=0;i< len(hours);i++  {
		if hours[i]>8 {
			arr = append(arr,arr[i]+1)
		}else{
			arr = append(arr,arr[i]-1)
		}
	}
	stack :=make([]int,1)
	for i:=1;i< len(arr);i++  {
		if arr[i]<arr[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	fmt.Println(arr,stack)
	ans:=0
	for i:= len(arr)-1;i>=0 ;i--  {
		for len(stack)>0 && arr[i]>arr[stack[len(stack)-1]] {
			fmt.Println(i-stack[len(stack)-1])
			ans = max(ans,i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return ans


}
func main() {
	fmt.Println(longestWPI([]int{
		9,9,6,0,6,6,9,



	}))
}
/*

	总结
	1. 我错了好多次啊，..因为太粗心了
	2. 我采取的策略是 从前到后扫描，如果>8则tmp+1，否则-1，如果tmp<0 就看看前面有没有tmp-1出现过，有就可以得到他们之间的长度是满足要求的，
	   但是要注意，tmp可能会出现两次相等的数值，那么我们最好的策略就是保留最前面的那个(代码上面有 标注1的地方)
		我采取的是哈希找前一个，其实还有二分的，我的复杂度是O（n）
	3. 官方有说是单调栈 0.0 我要看看
	4. 我用单调栈做出来了，但是理解还是不太透彻 0.0..
*/