package main
func numberOfSubarrays(nums []int, k int) int {
	even:=make([]int,1)
	oddIndex:=make([]int,0)
	for i:=1;i<=len(nums);i++{
		if nums[i-1]&1==1{
			oddIndex = append(oddIndex,i-1)
		}
		even = append(even,even[i-1] + int((nums[i-1]+1)&1))
	}
	ans:=0
	for i:=k-1;i<len(oddIndex);i++{
		begin:=oddIndex[i-k+1]
		end:=oddIndex[i]



		cntBegin:=i-k+1-1
		if cntBegin==-1{
			cntBegin=0
		}else{
			cntBegin = oddIndex[cntBegin]
		}
		cntEnd:=i+1
		if cntEnd==len(oddIndex){
			cntEnd = len(nums)
		}else{
			cntEnd = oddIndex[cntEnd]
		}

		ans += (even[begin] - even[cntBegin] + 1)  *(even[cntEnd] - even[end] + 1)

	}
	return ans
}
func main() {

}

/*
	总结
	1. 我AC的思想是  获取满足条件的最小子数组前后的偶数的个数，设为x,y
	   则对于这一组来说，总共可以构成(x+1)*(y+1)的优美数组
	2. 官方题解有很简单的AC法 0.0我要看看
	3. 这一题居然可以变化为两数和...

*/