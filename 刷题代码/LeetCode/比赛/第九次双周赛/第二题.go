package 第九次双周赛

import (
	"fmt"
	"sort"
)

func maxNumberOfApples(arr []int) int {
	sum:=5000
	sort.Ints(arr)
	ans:=0
	for i:=0;i< len(arr);i++  {
		sum-=arr[i]
		if sum>=0 {
			ans++
		}
	}
	return ans
}
func main() {
	fmt.Println(maxNumberOfApples([]int{

	}))
}
