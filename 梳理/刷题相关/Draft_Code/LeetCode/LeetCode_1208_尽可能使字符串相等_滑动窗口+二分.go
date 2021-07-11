package main


func equalSubstring(s string, t string, maxCost int) int {
	sum:=make([]int,1)
	for i:=1;i<=len(s);i++{
		sum = append(sum,sum[i-1] + abs(int(s[i-1])-int(t[i-1])))	// 构造
	}
	r:=0
	ans:=0
	for r<len(sum){
		index:=search1(sum,0,r,sum[r]-maxCost)	// 找到第一个≥sum[r]-maxCost的下标
		ans = max(r-index,ans)
		r++
	}
	return ans
}
// 第一个大于等于target的   -> lower_bound
func search1(A []int,l int,r int,target int) int{
	for l<=r{
		mid:=(l+r)/2
		if A[mid]==target{
			r = mid - 1
		}else{
			if A[mid]>target{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}
	}
	return l
}
// 最后一个小于等于target的
func search2(A []int,l int,r int,target int) int{
	for l<=r{
		mid:=(l+r)/2
		if A[mid]==target{
			l = mid + 1
		}else{
			if A[mid]>target{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}
	}
	return r
}


func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func abs(a int) int{
	if a>0{
		return int(a)
	}
	return int(-a)
}

/*
	总结
	1. 这代码的话不是左边一次缩短一格，而是通过二分查找的方式缩短 0.0
	2. lower_bound找的是第一个≥的，而upper_bound找的是第一个＞的
*/