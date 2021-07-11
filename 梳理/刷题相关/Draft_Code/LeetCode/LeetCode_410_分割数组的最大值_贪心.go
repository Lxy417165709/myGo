package main

import "fmt"

func count(nums []int,target int) int{
	tmp:=0
	ans:=1

	for i:=0;i<len(nums);i++{
		tmp+=nums[i]
		if tmp>target{
			tmp=nums[i]
			ans++
		}
	}

	return ans
}

func splitArray(nums []int, m int) int {
	l,r:=0,0

	for i:=0;i<len(nums);i++{
		l = max(l,nums[i])  // 很重要，l不能随便设
		r+=nums[i]
	}
	cnt:=1
	for l<r{
		mid := (l+r)/2

		cnt=count(nums,mid)
		fmt.Println(mid,cnt)
		if cnt>m{
			l = mid + 1 // 这个也和重要。。不然会死循环
		}else{
			r = mid
		}

	}
	return l

}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main() {

}
/*
	总结
	1.  这题有贪心性质，即尽可能的分出多个段，然后通过段数来判断
		首先分析题意，可以得出结论，结果必定落在【max（nums）， sum（nums）】这个区间内，
		因为左端点对应每个单独的元素构成一个子数组，右端点对应所有元素构成一个子数组。
		然后可以利用二分查找法逐步缩小区间范围，当区间长度为1时，即找到了最终答案。
		每次二分查找就是先算一个mid值，这个mid就是代表当前猜测的答案，
		然后模拟一下划分子数组的过程，可以得到用这个mid值会一共得到的子区间数cnt，然后比较cnt和m的关系，来更新区间范围。
*/